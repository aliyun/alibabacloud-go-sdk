// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAssetWatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateAssetWatchCmd) *CreateAssetWatchRequest
	GetBody() *CreateAssetWatchCmd
}

type CreateAssetWatchRequest struct {
	Body *CreateAssetWatchCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAssetWatchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAssetWatchRequest) GoString() string {
	return s.String()
}

func (s *CreateAssetWatchRequest) GetBody() *CreateAssetWatchCmd {
	return s.Body
}

func (s *CreateAssetWatchRequest) SetBody(v *CreateAssetWatchCmd) *CreateAssetWatchRequest {
	s.Body = v
	return s
}

func (s *CreateAssetWatchRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
