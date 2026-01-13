// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadApiCallDailyDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DownloadApiCallDailyDetailCmd) *DownloadApiCallDailyDetailRequest
	GetBody() *DownloadApiCallDailyDetailCmd
}

type DownloadApiCallDailyDetailRequest struct {
	Body *DownloadApiCallDailyDetailCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadApiCallDailyDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadApiCallDailyDetailRequest) GoString() string {
	return s.String()
}

func (s *DownloadApiCallDailyDetailRequest) GetBody() *DownloadApiCallDailyDetailCmd {
	return s.Body
}

func (s *DownloadApiCallDailyDetailRequest) SetBody(v *DownloadApiCallDailyDetailCmd) *DownloadApiCallDailyDetailRequest {
	s.Body = v
	return s
}

func (s *DownloadApiCallDailyDetailRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
