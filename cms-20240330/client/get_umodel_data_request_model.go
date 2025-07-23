// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUmodelDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v interface{}) *GetUmodelDataRequest
	GetContent() interface{}
	SetMethod(v string) *GetUmodelDataRequest
	GetMethod() *string
}

type GetUmodelDataRequest struct {
	// example:
	//
	// {
	//
	// 	"filter": {
	//
	// 		"domains": []
	//
	// 	},
	//
	// 	"offset": 0,
	//
	// 	"size": 100000
	//
	// }
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ListData
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
}

func (s GetUmodelDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUmodelDataRequest) GoString() string {
	return s.String()
}

func (s *GetUmodelDataRequest) GetContent() interface{} {
	return s.Content
}

func (s *GetUmodelDataRequest) GetMethod() *string {
	return s.Method
}

func (s *GetUmodelDataRequest) SetContent(v interface{}) *GetUmodelDataRequest {
	s.Content = v
	return s
}

func (s *GetUmodelDataRequest) SetMethod(v string) *GetUmodelDataRequest {
	s.Method = &v
	return s
}

func (s *GetUmodelDataRequest) Validate() error {
	return dara.Validate(s)
}
