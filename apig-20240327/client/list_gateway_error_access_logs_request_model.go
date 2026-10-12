// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayErrorAccessLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthority(v string) *ListGatewayErrorAccessLogsRequest
	GetAuthority() *string
	SetEndTime(v int64) *ListGatewayErrorAccessLogsRequest
	GetEndTime() *int64
	SetGatewayRequestId(v string) *ListGatewayErrorAccessLogsRequest
	GetGatewayRequestId() *string
	SetPath(v string) *ListGatewayErrorAccessLogsRequest
	GetPath() *string
	SetResponseCode(v string) *ListGatewayErrorAccessLogsRequest
	GetResponseCode() *string
	SetRouteName(v string) *ListGatewayErrorAccessLogsRequest
	GetRouteName() *string
	SetStartTime(v int64) *ListGatewayErrorAccessLogsRequest
	GetStartTime() *int64
}

type ListGatewayErrorAccessLogsRequest struct {
	// The domain name filter. This parameter performs an exact match against the authority field in Simple Log Service (SLS).
	//
	// example:
	//
	// api.example.com
	Authority *string `json:"authority,omitempty" xml:"authority,omitempty"`
	// The end of the query window, in UNIX seconds. The value cannot be later than the current time, and the time span cannot exceed 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1756348800
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The gateway request ID (x-acs-request-id) for exact match filtering.
	//
	// example:
	//
	// E5D4xxxx
	GatewayRequestId *string `json:"gatewayRequestId,omitempty" xml:"gatewayRequestId,omitempty"`
	// The request path filter.
	//
	// example:
	//
	// /api/v1/pets
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The HTTP error code. Only three-digit numbers in the 400–599 range are accepted. This parameter is required. If you do not specify this parameter, the API returns InvalidParameter.WithDetail (400).
	//
	// example:
	//
	// 502
	ResponseCode *string `json:"responseCode,omitempty" xml:"responseCode,omitempty"`
	// The route name filter.
	//
	// example:
	//
	// my-route
	RouteName *string `json:"routeName,omitempty" xml:"routeName,omitempty"`
	// The start of the query window, in UNIX seconds. The value must be no earlier than 24 hours before the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1756262400
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListGatewayErrorAccessLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayErrorAccessLogsRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayErrorAccessLogsRequest) GetAuthority() *string {
	return s.Authority
}

func (s *ListGatewayErrorAccessLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListGatewayErrorAccessLogsRequest) GetGatewayRequestId() *string {
	return s.GatewayRequestId
}

func (s *ListGatewayErrorAccessLogsRequest) GetPath() *string {
	return s.Path
}

func (s *ListGatewayErrorAccessLogsRequest) GetResponseCode() *string {
	return s.ResponseCode
}

func (s *ListGatewayErrorAccessLogsRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *ListGatewayErrorAccessLogsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListGatewayErrorAccessLogsRequest) SetAuthority(v string) *ListGatewayErrorAccessLogsRequest {
	s.Authority = &v
	return s
}

func (s *ListGatewayErrorAccessLogsRequest) SetEndTime(v int64) *ListGatewayErrorAccessLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListGatewayErrorAccessLogsRequest) SetGatewayRequestId(v string) *ListGatewayErrorAccessLogsRequest {
	s.GatewayRequestId = &v
	return s
}

func (s *ListGatewayErrorAccessLogsRequest) SetPath(v string) *ListGatewayErrorAccessLogsRequest {
	s.Path = &v
	return s
}

func (s *ListGatewayErrorAccessLogsRequest) SetResponseCode(v string) *ListGatewayErrorAccessLogsRequest {
	s.ResponseCode = &v
	return s
}

func (s *ListGatewayErrorAccessLogsRequest) SetRouteName(v string) *ListGatewayErrorAccessLogsRequest {
	s.RouteName = &v
	return s
}

func (s *ListGatewayErrorAccessLogsRequest) SetStartTime(v int64) *ListGatewayErrorAccessLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListGatewayErrorAccessLogsRequest) Validate() error {
	return dara.Validate(s)
}
