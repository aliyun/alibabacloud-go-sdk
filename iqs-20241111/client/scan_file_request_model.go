// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ScanFileInput) *ScanFileRequest
	GetBody() *ScanFileInput
}

type ScanFileRequest struct {
	Body *ScanFileInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScanFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ScanFileRequest) GoString() string {
	return s.String()
}

func (s *ScanFileRequest) GetBody() *ScanFileInput {
	return s.Body
}

func (s *ScanFileRequest) SetBody(v *ScanFileInput) *ScanFileRequest {
	s.Body = v
	return s
}

func (s *ScanFileRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
