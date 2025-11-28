// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllPrivacyAlgorithmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllPrivacyAlgorithmResponseBody
	GetCode() *string
	SetData(v []*ListAllPrivacyAlgorithmResponseBodyData) *ListAllPrivacyAlgorithmResponseBody
	GetData() []*ListAllPrivacyAlgorithmResponseBodyData
	SetHttpStatusCode(v int32) *ListAllPrivacyAlgorithmResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAllPrivacyAlgorithmResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllPrivacyAlgorithmResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllPrivacyAlgorithmResponseBody
	GetSuccess() *bool
}

type ListAllPrivacyAlgorithmResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListAllPrivacyAlgorithmResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAllPrivacyAlgorithmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllPrivacyAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllPrivacyAlgorithmResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllPrivacyAlgorithmResponseBody) GetData() []*ListAllPrivacyAlgorithmResponseBodyData {
	return s.Data
}

func (s *ListAllPrivacyAlgorithmResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAllPrivacyAlgorithmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllPrivacyAlgorithmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllPrivacyAlgorithmResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllPrivacyAlgorithmResponseBody) SetCode(v string) *ListAllPrivacyAlgorithmResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllPrivacyAlgorithmResponseBody) SetData(v []*ListAllPrivacyAlgorithmResponseBodyData) *ListAllPrivacyAlgorithmResponseBody {
	s.Data = v
	return s
}

func (s *ListAllPrivacyAlgorithmResponseBody) SetHttpStatusCode(v int32) *ListAllPrivacyAlgorithmResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAllPrivacyAlgorithmResponseBody) SetMessage(v string) *ListAllPrivacyAlgorithmResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllPrivacyAlgorithmResponseBody) SetRequestId(v string) *ListAllPrivacyAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllPrivacyAlgorithmResponseBody) SetSuccess(v bool) *ListAllPrivacyAlgorithmResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllPrivacyAlgorithmResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllPrivacyAlgorithmResponseBodyData struct {
	AlgImplList []*string `json:"AlgImplList,omitempty" xml:"AlgImplList,omitempty" type:"Repeated"`
	AlgType     *string   `json:"AlgType,omitempty" xml:"AlgType,omitempty"`
}

func (s ListAllPrivacyAlgorithmResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllPrivacyAlgorithmResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllPrivacyAlgorithmResponseBodyData) GetAlgImplList() []*string {
	return s.AlgImplList
}

func (s *ListAllPrivacyAlgorithmResponseBodyData) GetAlgType() *string {
	return s.AlgType
}

func (s *ListAllPrivacyAlgorithmResponseBodyData) SetAlgImplList(v []*string) *ListAllPrivacyAlgorithmResponseBodyData {
	s.AlgImplList = v
	return s
}

func (s *ListAllPrivacyAlgorithmResponseBodyData) SetAlgType(v string) *ListAllPrivacyAlgorithmResponseBodyData {
	s.AlgType = &v
	return s
}

func (s *ListAllPrivacyAlgorithmResponseBodyData) Validate() error {
	return dara.Validate(s)
}
