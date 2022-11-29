/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package producer

import (
	"net/http"

	"github.com/free5gc/nssf/internal/logger"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

// HandleNSSAIAvailabilityUnsubscribe - Deletes an already existing NSSAI availability notification subscription
func HandleNSSAIAvailabilityUnsubscribe(request *httpwrapper.Request) *httpwrapper.Response {
	logger.Nssaiavailability.Infof("Handle NSSAIAvailabilityUnsubscribe")

	subscriptionID := request.Params["subscriptionId"]

	problemDetails := NSSAIAvailabilityUnsubscribeProcedure(subscriptionID)

	if problemDetails == nil {
		return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
	} else if problemDetails != nil {
		return httpwrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	}
	problemDetails = &models.ProblemDetails{
		Status: http.StatusForbidden,
		Cause:  "UNSPECIFIED",
	}
	return httpwrapper.NewResponse(http.StatusForbidden, nil, problemDetails)
}
