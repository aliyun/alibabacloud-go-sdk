// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUuidValidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckUuidValidResponseBody
	GetCode() *string
	SetData(v *CheckUuidValidResponseBodyData) *CheckUuidValidResponseBody
	GetData() *CheckUuidValidResponseBodyData
	SetMessage(v string) *CheckUuidValidResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckUuidValidResponseBody
	GetRequestId() *string
}

type CheckUuidValidResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CheckUuidValidResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckUuidValidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckUuidValidResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUuidValidResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckUuidValidResponseBody) GetData() *CheckUuidValidResponseBodyData {
	return s.Data
}

func (s *CheckUuidValidResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckUuidValidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckUuidValidResponseBody) SetCode(v string) *CheckUuidValidResponseBody {
	s.Code = &v
	return s
}

func (s *CheckUuidValidResponseBody) SetData(v *CheckUuidValidResponseBodyData) *CheckUuidValidResponseBody {
	s.Data = v
	return s
}

func (s *CheckUuidValidResponseBody) SetMessage(v string) *CheckUuidValidResponseBody {
	s.Message = &v
	return s
}

func (s *CheckUuidValidResponseBody) SetRequestId(v string) *CheckUuidValidResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUuidValidResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckUuidValidResponseBodyData struct {
	NewUpgrade *bool `json:"NewUpgrade,omitempty" xml:"NewUpgrade,omitempty"`
}

func (s CheckUuidValidResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckUuidValidResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckUuidValidResponseBodyData) GetNewUpgrade() *bool {
	return s.NewUpgrade
}

func (s *CheckUuidValidResponseBodyData) SetNewUpgrade(v bool) *CheckUuidValidResponseBodyData {
	s.NewUpgrade = &v
	return s
}

func (s *CheckUuidValidResponseBodyData) Validate() error {
	return dara.Validate(s)
}
