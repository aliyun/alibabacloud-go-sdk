// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaskingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteMaskingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMaskingRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMaskingRulesResponseBody
	GetSuccess() *bool
}

type DeleteMaskingRulesResponseBody struct {
	// The response message.
	//
	// > If the request is successful, `Successful` is returned. If the request fails, an error message is returned, such as an error code.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2BCEE25B-797C-426B-BA7B-D28CCF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMaskingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaskingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaskingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMaskingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMaskingRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMaskingRulesResponseBody) SetMessage(v string) *DeleteMaskingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMaskingRulesResponseBody) SetRequestId(v string) *DeleteMaskingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaskingRulesResponseBody) SetSuccess(v bool) *DeleteMaskingRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMaskingRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
