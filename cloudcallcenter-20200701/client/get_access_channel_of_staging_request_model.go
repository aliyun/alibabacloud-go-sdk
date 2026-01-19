// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessChannelOfStagingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSearchPattern(v string) *GetAccessChannelOfStagingRequest
	GetSearchPattern() *string
}

type GetAccessChannelOfStagingRequest struct {
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s GetAccessChannelOfStagingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessChannelOfStagingRequest) GoString() string {
	return s.String()
}

func (s *GetAccessChannelOfStagingRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *GetAccessChannelOfStagingRequest) SetSearchPattern(v string) *GetAccessChannelOfStagingRequest {
	s.SearchPattern = &v
	return s
}

func (s *GetAccessChannelOfStagingRequest) Validate() error {
	return dara.Validate(s)
}
