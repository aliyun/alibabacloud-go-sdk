// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSAntChainConsortiumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBaaSAntChainConsortiumResponseBody
	GetCode() *string
	SetData(v []*ListBaaSAntChainConsortiumResponseBodyData) *ListBaaSAntChainConsortiumResponseBody
	GetData() []*ListBaaSAntChainConsortiumResponseBodyData
	SetHttpStatusCode(v int32) *ListBaaSAntChainConsortiumResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBaaSAntChainConsortiumResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBaaSAntChainConsortiumResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBaaSAntChainConsortiumResponseBody
	GetSuccess() *bool
}

type ListBaaSAntChainConsortiumResponseBody struct {
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListBaaSAntChainConsortiumResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBaaSAntChainConsortiumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSAntChainConsortiumResponseBody) GoString() string {
	return s.String()
}

func (s *ListBaaSAntChainConsortiumResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBaaSAntChainConsortiumResponseBody) GetData() []*ListBaaSAntChainConsortiumResponseBodyData {
	return s.Data
}

func (s *ListBaaSAntChainConsortiumResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBaaSAntChainConsortiumResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBaaSAntChainConsortiumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBaaSAntChainConsortiumResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBaaSAntChainConsortiumResponseBody) SetCode(v string) *ListBaaSAntChainConsortiumResponseBody {
	s.Code = &v
	return s
}

func (s *ListBaaSAntChainConsortiumResponseBody) SetData(v []*ListBaaSAntChainConsortiumResponseBodyData) *ListBaaSAntChainConsortiumResponseBody {
	s.Data = v
	return s
}

func (s *ListBaaSAntChainConsortiumResponseBody) SetHttpStatusCode(v int32) *ListBaaSAntChainConsortiumResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBaaSAntChainConsortiumResponseBody) SetMessage(v string) *ListBaaSAntChainConsortiumResponseBody {
	s.Message = &v
	return s
}

func (s *ListBaaSAntChainConsortiumResponseBody) SetRequestId(v string) *ListBaaSAntChainConsortiumResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBaaSAntChainConsortiumResponseBody) SetSuccess(v bool) *ListBaaSAntChainConsortiumResponseBody {
	s.Success = &v
	return s
}

func (s *ListBaaSAntChainConsortiumResponseBody) Validate() error {
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

type ListBaaSAntChainConsortiumResponseBodyData struct {
	BaaSAntChainConsortiumId   *string `json:"BaaSAntChainConsortiumId,omitempty" xml:"BaaSAntChainConsortiumId,omitempty"`
	BaaSAntChainConsortiumName *string `json:"BaaSAntChainConsortiumName,omitempty" xml:"BaaSAntChainConsortiumName,omitempty"`
}

func (s ListBaaSAntChainConsortiumResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSAntChainConsortiumResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBaaSAntChainConsortiumResponseBodyData) GetBaaSAntChainConsortiumId() *string {
	return s.BaaSAntChainConsortiumId
}

func (s *ListBaaSAntChainConsortiumResponseBodyData) GetBaaSAntChainConsortiumName() *string {
	return s.BaaSAntChainConsortiumName
}

func (s *ListBaaSAntChainConsortiumResponseBodyData) SetBaaSAntChainConsortiumId(v string) *ListBaaSAntChainConsortiumResponseBodyData {
	s.BaaSAntChainConsortiumId = &v
	return s
}

func (s *ListBaaSAntChainConsortiumResponseBodyData) SetBaaSAntChainConsortiumName(v string) *ListBaaSAntChainConsortiumResponseBodyData {
	s.BaaSAntChainConsortiumName = &v
	return s
}

func (s *ListBaaSAntChainConsortiumResponseBodyData) Validate() error {
	return dara.Validate(s)
}
