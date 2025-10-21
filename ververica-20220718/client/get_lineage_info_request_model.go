// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLineageInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *GetLineageInfoParams) *GetLineageInfoRequest
	GetBody() *GetLineageInfoParams
}

type GetLineageInfoRequest struct {
	// The parameters about the lineage information.
	Body *GetLineageInfoParams `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLineageInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLineageInfoRequest) GoString() string {
	return s.String()
}

func (s *GetLineageInfoRequest) GetBody() *GetLineageInfoParams {
	return s.Body
}

func (s *GetLineageInfoRequest) SetBody(v *GetLineageInfoParams) *GetLineageInfoRequest {
	s.Body = v
	return s
}

func (s *GetLineageInfoRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
