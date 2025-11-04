// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotwordLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotwordLibraryId(v string) *GetHotwordLibraryRequest
	GetHotwordLibraryId() *string
}

type GetHotwordLibraryRequest struct {
	// The ID of the hotword library.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	HotwordLibraryId *string `json:"HotwordLibraryId,omitempty" xml:"HotwordLibraryId,omitempty"`
}

func (s GetHotwordLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotwordLibraryRequest) GoString() string {
	return s.String()
}

func (s *GetHotwordLibraryRequest) GetHotwordLibraryId() *string {
	return s.HotwordLibraryId
}

func (s *GetHotwordLibraryRequest) SetHotwordLibraryId(v string) *GetHotwordLibraryRequest {
	s.HotwordLibraryId = &v
	return s
}

func (s *GetHotwordLibraryRequest) Validate() error {
	return dara.Validate(s)
}
