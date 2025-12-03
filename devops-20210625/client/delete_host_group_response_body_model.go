// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteHostGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteHostGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteHostGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHostGroupResponseBody
	GetSuccess() *bool
}

type DeleteHostGroupResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteHostGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHostGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteHostGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteHostGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHostGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHostGroupResponseBody) SetErrorCode(v string) *DeleteHostGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteHostGroupResponseBody) SetErrorMessage(v string) *DeleteHostGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteHostGroupResponseBody) SetRequestId(v string) *DeleteHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHostGroupResponseBody) SetSuccess(v bool) *DeleteHostGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHostGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
