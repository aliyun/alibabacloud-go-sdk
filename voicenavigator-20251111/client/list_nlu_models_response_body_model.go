// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNluModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListNluModelsResponseBody
	GetCode() *string
	SetData(v []*ListNluModelsResponseBodyData) *ListNluModelsResponseBody
	GetData() []*ListNluModelsResponseBodyData
	SetHttpStatusCode(v int32) *ListNluModelsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListNluModelsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListNluModelsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListNluModelsResponseBody
	GetRequestId() *string
}

type ListNluModelsResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListNluModelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 7401D698-0AAC-5909-B68E-88C68805FFB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNluModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNluModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNluModelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListNluModelsResponseBody) GetData() []*ListNluModelsResponseBodyData {
	return s.Data
}

func (s *ListNluModelsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListNluModelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNluModelsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListNluModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNluModelsResponseBody) SetCode(v string) *ListNluModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListNluModelsResponseBody) SetData(v []*ListNluModelsResponseBodyData) *ListNluModelsResponseBody {
	s.Data = v
	return s
}

func (s *ListNluModelsResponseBody) SetHttpStatusCode(v int32) *ListNluModelsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNluModelsResponseBody) SetMessage(v string) *ListNluModelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListNluModelsResponseBody) SetParams(v []*string) *ListNluModelsResponseBody {
	s.Params = v
	return s
}

func (s *ListNluModelsResponseBody) SetRequestId(v string) *ListNluModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNluModelsResponseBody) Validate() error {
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

type ListNluModelsResponseBodyData struct {
	// example:
	//
	// qwen-plus
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListNluModelsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNluModelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNluModelsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListNluModelsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListNluModelsResponseBodyData) SetId(v string) *ListNluModelsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListNluModelsResponseBodyData) SetName(v string) *ListNluModelsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListNluModelsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
