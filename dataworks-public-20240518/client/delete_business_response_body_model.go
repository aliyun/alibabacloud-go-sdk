// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBusinessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteBusinessResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteBusinessResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteBusinessResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteBusinessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBusinessResponseBody
	GetSuccess() *bool
}

type DeleteBusinessResponseBody struct {
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBusinessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBusinessResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBusinessResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteBusinessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteBusinessResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBusinessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBusinessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBusinessResponseBody) SetErrorCode(v string) *DeleteBusinessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteBusinessResponseBody) SetErrorMessage(v string) *DeleteBusinessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteBusinessResponseBody) SetHttpStatusCode(v int32) *DeleteBusinessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBusinessResponseBody) SetRequestId(v string) *DeleteBusinessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBusinessResponseBody) SetSuccess(v bool) *DeleteBusinessResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBusinessResponseBody) Validate() error {
	return dara.Validate(s)
}
