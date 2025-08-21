// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSDGId(v string) *SaveSDGRequest
	GetSDGId() *string
}

type SaveSDGRequest struct {
	// The ID of the SDG to be saved.
	//
	// example:
	//
	// sdg-xxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s SaveSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSDGRequest) GoString() string {
	return s.String()
}

func (s *SaveSDGRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *SaveSDGRequest) SetSDGId(v string) *SaveSDGRequest {
	s.SDGId = &v
	return s
}

func (s *SaveSDGRequest) Validate() error {
	return dara.Validate(s)
}
