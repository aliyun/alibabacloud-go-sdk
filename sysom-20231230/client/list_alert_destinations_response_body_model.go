// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertDestinationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAlertDestinationsResponseBody
	GetCode() *string
	SetData(v interface{}) *ListAlertDestinationsResponseBody
	GetData() interface{}
	SetMaxResults(v int32) *ListAlertDestinationsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAlertDestinationsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAlertDestinationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAlertDestinationsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListAlertDestinationsResponseBody
	GetTotal() *int32
}

type ListAlertDestinationsResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// server error
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// c2f78a783f49457caba6bace6f6f79e4
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 623
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAlertDestinationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertDestinationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAlertDestinationsResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListAlertDestinationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertDestinationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlertDestinationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertDestinationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertDestinationsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAlertDestinationsResponseBody) SetCode(v string) *ListAlertDestinationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlertDestinationsResponseBody) SetData(v interface{}) *ListAlertDestinationsResponseBody {
	s.Data = v
	return s
}

func (s *ListAlertDestinationsResponseBody) SetMaxResults(v int32) *ListAlertDestinationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAlertDestinationsResponseBody) SetMessage(v string) *ListAlertDestinationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlertDestinationsResponseBody) SetNextToken(v string) *ListAlertDestinationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAlertDestinationsResponseBody) SetRequestId(v string) *ListAlertDestinationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertDestinationsResponseBody) SetTotal(v int32) *ListAlertDestinationsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAlertDestinationsResponseBody) Validate() error {
	return dara.Validate(s)
}
