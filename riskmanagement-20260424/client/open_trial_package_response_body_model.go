// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenTrialPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OpenTrialPackageResponseBody
	GetCode() *string
	SetData(v *OpenTrialPackageResponseBodyData) *OpenTrialPackageResponseBody
	GetData() *OpenTrialPackageResponseBodyData
	SetMessage(v string) *OpenTrialPackageResponseBody
	GetMessage() *string
	SetRequestId(v string) *OpenTrialPackageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OpenTrialPackageResponseBody
	GetSuccess() *bool
}

type OpenTrialPackageResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *OpenTrialPackageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OpenTrialPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenTrialPackageResponseBody) GoString() string {
	return s.String()
}

func (s *OpenTrialPackageResponseBody) GetCode() *string {
	return s.Code
}

func (s *OpenTrialPackageResponseBody) GetData() *OpenTrialPackageResponseBodyData {
	return s.Data
}

func (s *OpenTrialPackageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenTrialPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenTrialPackageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OpenTrialPackageResponseBody) SetCode(v string) *OpenTrialPackageResponseBody {
	s.Code = &v
	return s
}

func (s *OpenTrialPackageResponseBody) SetData(v *OpenTrialPackageResponseBodyData) *OpenTrialPackageResponseBody {
	s.Data = v
	return s
}

func (s *OpenTrialPackageResponseBody) SetMessage(v string) *OpenTrialPackageResponseBody {
	s.Message = &v
	return s
}

func (s *OpenTrialPackageResponseBody) SetRequestId(v string) *OpenTrialPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenTrialPackageResponseBody) SetSuccess(v bool) *OpenTrialPackageResponseBody {
	s.Success = &v
	return s
}

func (s *OpenTrialPackageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OpenTrialPackageResponseBodyData struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenTrialPackageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OpenTrialPackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *OpenTrialPackageResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenTrialPackageResponseBodyData) SetRequestId(v string) *OpenTrialPackageResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *OpenTrialPackageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
