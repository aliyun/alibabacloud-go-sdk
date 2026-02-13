// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceUpdateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AISearchResourceUpdateResponseBodyData) *AISearchResourceUpdateResponseBody
	GetData() *AISearchResourceUpdateResponseBodyData
	SetRequestId(v string) *AISearchResourceUpdateResponseBody
	GetRequestId() *string
}

type AISearchResourceUpdateResponseBody struct {
	Data *AISearchResourceUpdateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// FDE51A3B-09C0-57E5-96FC-31E85EEFF318
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AISearchResourceUpdateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *AISearchResourceUpdateResponseBody) GetData() *AISearchResourceUpdateResponseBodyData {
	return s.Data
}

func (s *AISearchResourceUpdateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AISearchResourceUpdateResponseBody) SetData(v *AISearchResourceUpdateResponseBodyData) *AISearchResourceUpdateResponseBody {
	s.Data = v
	return s
}

func (s *AISearchResourceUpdateResponseBody) SetRequestId(v string) *AISearchResourceUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AISearchResourceUpdateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AISearchResourceUpdateResponseBodyData struct {
	// example:
	//
	// WzMGQZwB7nQEs3Qk3ajH
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// miniapp
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AISearchResourceUpdateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceUpdateResponseBodyData) GoString() string {
	return s.String()
}

func (s *AISearchResourceUpdateResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *AISearchResourceUpdateResponseBodyData) GetType() *string {
	return s.Type
}

func (s *AISearchResourceUpdateResponseBodyData) SetResourceId(v string) *AISearchResourceUpdateResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *AISearchResourceUpdateResponseBodyData) SetType(v string) *AISearchResourceUpdateResponseBodyData {
	s.Type = &v
	return s
}

func (s *AISearchResourceUpdateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
