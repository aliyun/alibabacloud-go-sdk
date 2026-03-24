// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ScaleServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ScaleServiceResponseBody
	GetRequestId() *string
}

type ScaleServiceResponseBody struct {
	// example:
	//
	// Service [foo] in region [cn-shanghai] is scaling
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScaleServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ScaleServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleServiceResponseBody) SetMessage(v string) *ScaleServiceResponseBody {
	s.Message = &v
	return s
}

func (s *ScaleServiceResponseBody) SetRequestId(v string) *ScaleServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
