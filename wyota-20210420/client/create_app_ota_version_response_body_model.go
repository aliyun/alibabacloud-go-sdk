// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppOtaVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAppOtaVersionResponseBody
	GetCode() *string
	SetData(v *CreateAppOtaVersionResponseBodyData) *CreateAppOtaVersionResponseBody
	GetData() *CreateAppOtaVersionResponseBodyData
	SetMessage(v string) *CreateAppOtaVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAppOtaVersionResponseBody
	GetRequestId() *string
}

type CreateAppOtaVersionResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateAppOtaVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppOtaVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppOtaVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppOtaVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAppOtaVersionResponseBody) GetData() *CreateAppOtaVersionResponseBodyData {
	return s.Data
}

func (s *CreateAppOtaVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAppOtaVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppOtaVersionResponseBody) SetCode(v string) *CreateAppOtaVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppOtaVersionResponseBody) SetData(v *CreateAppOtaVersionResponseBodyData) *CreateAppOtaVersionResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppOtaVersionResponseBody) SetMessage(v string) *CreateAppOtaVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppOtaVersionResponseBody) SetRequestId(v string) *CreateAppOtaVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppOtaVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAppOtaVersionResponseBodyData struct {
	VersionUid *string `json:"VersionUid,omitempty" xml:"VersionUid,omitempty"`
}

func (s CreateAppOtaVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAppOtaVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppOtaVersionResponseBodyData) GetVersionUid() *string {
	return s.VersionUid
}

func (s *CreateAppOtaVersionResponseBodyData) SetVersionUid(v string) *CreateAppOtaVersionResponseBodyData {
	s.VersionUid = &v
	return s
}

func (s *CreateAppOtaVersionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
