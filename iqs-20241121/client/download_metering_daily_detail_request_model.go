// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadMeteringDailyDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DownloadMeteringDailyDetailCmd) *DownloadMeteringDailyDetailRequest
	GetBody() *DownloadMeteringDailyDetailCmd
}

type DownloadMeteringDailyDetailRequest struct {
	Body *DownloadMeteringDailyDetailCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadMeteringDailyDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadMeteringDailyDetailRequest) GoString() string {
	return s.String()
}

func (s *DownloadMeteringDailyDetailRequest) GetBody() *DownloadMeteringDailyDetailCmd {
	return s.Body
}

func (s *DownloadMeteringDailyDetailRequest) SetBody(v *DownloadMeteringDailyDetailCmd) *DownloadMeteringDailyDetailRequest {
	s.Body = v
	return s
}

func (s *DownloadMeteringDailyDetailRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
