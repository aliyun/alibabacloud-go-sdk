// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AISearchResourceDeleteResponseBodyData) *AISearchResourceDeleteResponseBody
	GetData() *AISearchResourceDeleteResponseBodyData
	SetRequestId(v string) *AISearchResourceDeleteResponseBody
	GetRequestId() *string
}

type AISearchResourceDeleteResponseBody struct {
	Data *AISearchResourceDeleteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AISearchResourceDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *AISearchResourceDeleteResponseBody) GetData() *AISearchResourceDeleteResponseBodyData {
	return s.Data
}

func (s *AISearchResourceDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AISearchResourceDeleteResponseBody) SetData(v *AISearchResourceDeleteResponseBodyData) *AISearchResourceDeleteResponseBody {
	s.Data = v
	return s
}

func (s *AISearchResourceDeleteResponseBody) SetRequestId(v string) *AISearchResourceDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *AISearchResourceDeleteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AISearchResourceDeleteResponseBodyData struct {
	// example:
	//
	// WzMGQZwB7nQEs3Qk3ajH
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// miniapp
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AISearchResourceDeleteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceDeleteResponseBodyData) GoString() string {
	return s.String()
}

func (s *AISearchResourceDeleteResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *AISearchResourceDeleteResponseBodyData) GetType() *string {
	return s.Type
}

func (s *AISearchResourceDeleteResponseBodyData) SetResourceId(v string) *AISearchResourceDeleteResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *AISearchResourceDeleteResponseBodyData) SetType(v string) *AISearchResourceDeleteResponseBodyData {
	s.Type = &v
	return s
}

func (s *AISearchResourceDeleteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
