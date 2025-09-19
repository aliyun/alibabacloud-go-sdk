// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageEventOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetImageEventOperationRequest
	GetId() *int64
	SetLang(v string) *GetImageEventOperationRequest
	GetLang() *string
}

type GetImageEventOperationRequest struct {
	// The primary key of the alert handling rule.
	//
	// example:
	//
	// 814163
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetImageEventOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageEventOperationRequest) GoString() string {
	return s.String()
}

func (s *GetImageEventOperationRequest) GetId() *int64 {
	return s.Id
}

func (s *GetImageEventOperationRequest) GetLang() *string {
	return s.Lang
}

func (s *GetImageEventOperationRequest) SetId(v int64) *GetImageEventOperationRequest {
	s.Id = &v
	return s
}

func (s *GetImageEventOperationRequest) SetLang(v string) *GetImageEventOperationRequest {
	s.Lang = &v
	return s
}

func (s *GetImageEventOperationRequest) Validate() error {
	return dara.Validate(s)
}
