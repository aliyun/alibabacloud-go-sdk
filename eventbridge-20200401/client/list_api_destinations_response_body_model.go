// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiDestinationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApiDestinationsResponseBody
	GetCode() *string
	SetData(v *ListApiDestinationsResponseBodyData) *ListApiDestinationsResponseBody
	GetData() *ListApiDestinationsResponseBodyData
	SetMessage(v string) *ListApiDestinationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApiDestinationsResponseBody
	GetRequestId() *string
}

type ListApiDestinationsResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListApiDestinationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message. If the request is successful, success is returned. If the request failed, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 96D7C0AB-DCE5-5E82-96B8-4725E1706BB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApiDestinationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiDestinationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApiDestinationsResponseBody) GetData() *ListApiDestinationsResponseBodyData {
	return s.Data
}

func (s *ListApiDestinationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApiDestinationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiDestinationsResponseBody) SetCode(v string) *ListApiDestinationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListApiDestinationsResponseBody) SetData(v *ListApiDestinationsResponseBodyData) *ListApiDestinationsResponseBody {
	s.Data = v
	return s
}

func (s *ListApiDestinationsResponseBody) SetMessage(v string) *ListApiDestinationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListApiDestinationsResponseBody) SetRequestId(v string) *ListApiDestinationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiDestinationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApiDestinationsResponseBodyData struct {
	// The API destinations.
	ApiDestinations []*ListApiDestinationsResponseBodyDataApiDestinations `json:"ApiDestinations,omitempty" xml:"ApiDestinations,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *float32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If excess return values exist, this parameter is returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *float32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListApiDestinationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListApiDestinationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApiDestinationsResponseBodyData) GetApiDestinations() []*ListApiDestinationsResponseBodyDataApiDestinations {
	return s.ApiDestinations
}

func (s *ListApiDestinationsResponseBodyData) GetMaxResults() *float32 {
	return s.MaxResults
}

func (s *ListApiDestinationsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiDestinationsResponseBodyData) GetTotal() *float32 {
	return s.Total
}

func (s *ListApiDestinationsResponseBodyData) SetApiDestinations(v []*ListApiDestinationsResponseBodyDataApiDestinations) *ListApiDestinationsResponseBodyData {
	s.ApiDestinations = v
	return s
}

func (s *ListApiDestinationsResponseBodyData) SetMaxResults(v float32) *ListApiDestinationsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListApiDestinationsResponseBodyData) SetNextToken(v string) *ListApiDestinationsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListApiDestinationsResponseBodyData) SetTotal(v float32) *ListApiDestinationsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListApiDestinationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListApiDestinationsResponseBodyDataApiDestinations struct {
	// The name of the API destination.
	//
	// example:
	//
	// api-destination-2
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	// The connection name.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the connection.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the API destination was created.
	//
	// example:
	//
	// 1665223213000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The request parameters that are configured for the API destination.
	HttpApiParameters *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty" type:"Struct"`
}

func (s ListApiDestinationsResponseBodyDataApiDestinations) String() string {
	return dara.Prettify(s)
}

func (s ListApiDestinationsResponseBodyDataApiDestinations) GoString() string {
	return s.String()
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) GetApiDestinationName() *string {
	return s.ApiDestinationName
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) GetDescription() *string {
	return s.Description
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) GetHttpApiParameters() *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters {
	return s.HttpApiParameters
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) SetApiDestinationName(v string) *ListApiDestinationsResponseBodyDataApiDestinations {
	s.ApiDestinationName = &v
	return s
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) SetConnectionName(v string) *ListApiDestinationsResponseBodyDataApiDestinations {
	s.ConnectionName = &v
	return s
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) SetDescription(v string) *ListApiDestinationsResponseBodyDataApiDestinations {
	s.Description = &v
	return s
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) SetGmtCreate(v int64) *ListApiDestinationsResponseBodyDataApiDestinations {
	s.GmtCreate = &v
	return s
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) SetHttpApiParameters(v *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) *ListApiDestinationsResponseBodyDataApiDestinations {
	s.HttpApiParameters = v
	return s
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) Validate() error {
	return dara.Validate(s)
}

type ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters struct {
	// The endpoint of the API destination.
	//
	// example:
	//
	// http://127.0.0.1:8001/api
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The HTTP request method. Valid values:
	//
	// - POST
	//
	// - GET
	//
	// - DELETE
	//
	// - PUT
	//
	// - HEAD
	//
	// - TRACE
	//
	// - PATCH
	//
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) String() string {
	return dara.Prettify(s)
}

func (s ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) GoString() string {
	return s.String()
}

func (s *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) GetMethod() *string {
	return s.Method
}

func (s *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) SetEndpoint(v string) *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters {
	s.Endpoint = &v
	return s
}

func (s *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) SetMethod(v string) *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters {
	s.Method = &v
	return s
}

func (s *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) Validate() error {
	return dara.Validate(s)
}
