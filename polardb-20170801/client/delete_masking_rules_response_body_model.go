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
	// The message that is returned for the request.
	//
	// > If the request is successful, `Successful` is returned. If the request fails, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2BCEE25B-797C-426B-BA7B-D28CCF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid value:
	//
	// 	- **true**
	//
	// 	- **false**
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
