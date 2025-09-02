// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLineageRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteLineageRelationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteLineageRelationResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteLineageRelationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteLineageRelationResponseBody
	GetRequestId() *string
	SetStatus(v bool) *DeleteLineageRelationResponseBody
	GetStatus() *bool
	SetSuccess(v bool) *DeleteLineageRelationResponseBody
	GetSuccess() *bool
}

type DeleteLineageRelationResponseBody struct {
	// Error code
	//
	// example:
	//
	// 1010040007
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Error message
	//
	// example:
	//
	// qualifiedName should be in format as entity-table.entity-guid
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Request ID: used for locating logs and troubleshooting
	//
	// example:
	//
	// 64B-587A-8CED-969E1973887FXXX-TT
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Operation result:
	//
	// true: Success
	//
	// false: Failure
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// Whether the call was successful. Values are as follows:
	//
	// true: success
	//
	// false: failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLineageRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLineageRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLineageRelationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteLineageRelationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteLineageRelationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteLineageRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLineageRelationResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *DeleteLineageRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteLineageRelationResponseBody) SetErrorCode(v string) *DeleteLineageRelationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteLineageRelationResponseBody) SetErrorMessage(v string) *DeleteLineageRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteLineageRelationResponseBody) SetHttpStatusCode(v int32) *DeleteLineageRelationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteLineageRelationResponseBody) SetRequestId(v string) *DeleteLineageRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLineageRelationResponseBody) SetStatus(v bool) *DeleteLineageRelationResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteLineageRelationResponseBody) SetSuccess(v bool) *DeleteLineageRelationResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLineageRelationResponseBody) Validate() error {
	return dara.Validate(s)
}
