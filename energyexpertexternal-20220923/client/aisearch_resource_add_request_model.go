// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceAddRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *AISearchResourceAddRequest
	GetData() interface{}
	SetType(v string) *AISearchResourceAddRequest
	GetType() *string
}

type AISearchResourceAddRequest struct {
	// example:
	//
	// {
	//
	//  "miniapp_id": "test_miniapp",
	//
	//       "title": "testApp",
	//
	//       "version": "1.0.5",
	//
	//       "description": "description-mock",
	//
	//       "slogan": "slogan-mock",
	//
	//       "icon": "https://img.alicdn.com/test_icon.png",
	//
	//       "detail_desc": "detail-mock"
	//
	// }
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// miniapp
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AISearchResourceAddRequest) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceAddRequest) GoString() string {
	return s.String()
}

func (s *AISearchResourceAddRequest) GetData() interface{} {
	return s.Data
}

func (s *AISearchResourceAddRequest) GetType() *string {
	return s.Type
}

func (s *AISearchResourceAddRequest) SetData(v interface{}) *AISearchResourceAddRequest {
	s.Data = v
	return s
}

func (s *AISearchResourceAddRequest) SetType(v string) *AISearchResourceAddRequest {
	s.Type = &v
	return s
}

func (s *AISearchResourceAddRequest) Validate() error {
	return dara.Validate(s)
}
