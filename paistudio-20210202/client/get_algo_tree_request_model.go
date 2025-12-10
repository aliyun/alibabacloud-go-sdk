// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgoTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *GetAlgoTreeRequest
	GetSource() *string
}

type GetAlgoTreeRequest struct {
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetAlgoTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlgoTreeRequest) GoString() string {
	return s.String()
}

func (s *GetAlgoTreeRequest) GetSource() *string {
	return s.Source
}

func (s *GetAlgoTreeRequest) SetSource(v string) *GetAlgoTreeRequest {
	s.Source = &v
	return s
}

func (s *GetAlgoTreeRequest) Validate() error {
	return dara.Validate(s)
}
