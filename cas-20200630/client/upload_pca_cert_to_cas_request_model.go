// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPcaCertToCasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *UploadPcaCertToCasRequest
	GetIds() *string
}

type UploadPcaCertToCasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 59425,59426
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s UploadPcaCertToCasRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadPcaCertToCasRequest) GoString() string {
	return s.String()
}

func (s *UploadPcaCertToCasRequest) GetIds() *string {
	return s.Ids
}

func (s *UploadPcaCertToCasRequest) SetIds(v string) *UploadPcaCertToCasRequest {
	s.Ids = &v
	return s
}

func (s *UploadPcaCertToCasRequest) Validate() error {
	return dara.Validate(s)
}
