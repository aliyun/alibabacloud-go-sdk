// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v bool) *GetImageRequest
	GetVerbose() *bool
}

type GetImageRequest struct {
	// Specifies whether to display non-essential information. Non-essential information currently includes labels.
	//
	// Valid values:
	//
	// - false (default): Does not display non-essential information.
	//
	// - true: Displays non-essential information.
	//
	// example:
	//
	// false
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetImageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetImageRequest) SetVerbose(v bool) *GetImageRequest {
	s.Verbose = &v
	return s
}

func (s *GetImageRequest) Validate() error {
	return dara.Validate(s)
}
