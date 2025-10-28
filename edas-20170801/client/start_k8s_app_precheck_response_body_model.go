// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartK8sAppPrecheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *StartK8sAppPrecheckResponseBody
	GetCode() *int32
	SetData(v *StartK8sAppPrecheckResponseBodyData) *StartK8sAppPrecheckResponseBody
	GetData() *StartK8sAppPrecheckResponseBodyData
	SetMessage(v string) *StartK8sAppPrecheckResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartK8sAppPrecheckResponseBody
	GetRequestId() *string
}

type StartK8sAppPrecheckResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *StartK8sAppPrecheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7638276F-****-****-884F-54CC0BC84A8D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartK8sAppPrecheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartK8sAppPrecheckResponseBody) GoString() string {
	return s.String()
}

func (s *StartK8sAppPrecheckResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *StartK8sAppPrecheckResponseBody) GetData() *StartK8sAppPrecheckResponseBodyData {
	return s.Data
}

func (s *StartK8sAppPrecheckResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartK8sAppPrecheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartK8sAppPrecheckResponseBody) SetCode(v int32) *StartK8sAppPrecheckResponseBody {
	s.Code = &v
	return s
}

func (s *StartK8sAppPrecheckResponseBody) SetData(v *StartK8sAppPrecheckResponseBodyData) *StartK8sAppPrecheckResponseBody {
	s.Data = v
	return s
}

func (s *StartK8sAppPrecheckResponseBody) SetMessage(v string) *StartK8sAppPrecheckResponseBody {
	s.Message = &v
	return s
}

func (s *StartK8sAppPrecheckResponseBody) SetRequestId(v string) *StartK8sAppPrecheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartK8sAppPrecheckResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartK8sAppPrecheckResponseBodyData struct {
	// The jobs and the details about the jobs.
	Jobs []*string `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
}

func (s StartK8sAppPrecheckResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartK8sAppPrecheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartK8sAppPrecheckResponseBodyData) GetJobs() []*string {
	return s.Jobs
}

func (s *StartK8sAppPrecheckResponseBodyData) SetJobs(v []*string) *StartK8sAppPrecheckResponseBodyData {
	s.Jobs = v
	return s
}

func (s *StartK8sAppPrecheckResponseBodyData) Validate() error {
	return dara.Validate(s)
}
