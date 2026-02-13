// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceAddResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AISearchResourceAddResponseBodyData) *AISearchResourceAddResponseBody
	GetData() *AISearchResourceAddResponseBodyData
	SetRequestId(v string) *AISearchResourceAddResponseBody
	GetRequestId() *string
}

type AISearchResourceAddResponseBody struct {
	Data *AISearchResourceAddResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 4A0AEC56-5C9A-5D47-93DF-7227836FFF82
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AISearchResourceAddResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceAddResponseBody) GoString() string {
	return s.String()
}

func (s *AISearchResourceAddResponseBody) GetData() *AISearchResourceAddResponseBodyData {
	return s.Data
}

func (s *AISearchResourceAddResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AISearchResourceAddResponseBody) SetData(v *AISearchResourceAddResponseBodyData) *AISearchResourceAddResponseBody {
	s.Data = v
	return s
}

func (s *AISearchResourceAddResponseBody) SetRequestId(v string) *AISearchResourceAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *AISearchResourceAddResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AISearchResourceAddResponseBodyData struct {
	// example:
	//
	// WzMGQZwB7nQEs3Qk3ajH
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// miniapp
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AISearchResourceAddResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceAddResponseBodyData) GoString() string {
	return s.String()
}

func (s *AISearchResourceAddResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *AISearchResourceAddResponseBodyData) GetType() *string {
	return s.Type
}

func (s *AISearchResourceAddResponseBodyData) SetResourceId(v string) *AISearchResourceAddResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *AISearchResourceAddResponseBodyData) SetType(v string) *AISearchResourceAddResponseBodyData {
	s.Type = &v
	return s
}

func (s *AISearchResourceAddResponseBodyData) Validate() error {
	return dara.Validate(s)
}
