// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvanceConfigFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetAdvanceConfigFileRequest
	GetFileName() *string
}

type GetAdvanceConfigFileRequest struct {
	// The name of the file
	//
	// This parameter is required.
	//
	// example:
	//
	// /intervene_dict/chn_ecommerce_general.dict
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetAdvanceConfigFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAdvanceConfigFileRequest) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetAdvanceConfigFileRequest) SetFileName(v string) *GetAdvanceConfigFileRequest {
	s.FileName = &v
	return s
}

func (s *GetAdvanceConfigFileRequest) Validate() error {
	return dara.Validate(s)
}
