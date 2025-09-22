// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncSignInInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v string) *SyncSignInInfoRequest
	GetActivityId() *string
	SetQRCode(v string) *SyncSignInInfoRequest
	GetQRCode() *string
}

type SyncSignInInfoRequest struct {
	// This parameter is required.
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// This parameter is required.
	QRCode *string `json:"QRCode,omitempty" xml:"QRCode,omitempty"`
}

func (s SyncSignInInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncSignInInfoRequest) GoString() string {
	return s.String()
}

func (s *SyncSignInInfoRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *SyncSignInInfoRequest) GetQRCode() *string {
	return s.QRCode
}

func (s *SyncSignInInfoRequest) SetActivityId(v string) *SyncSignInInfoRequest {
	s.ActivityId = &v
	return s
}

func (s *SyncSignInInfoRequest) SetQRCode(v string) *SyncSignInInfoRequest {
	s.QRCode = &v
	return s
}

func (s *SyncSignInInfoRequest) Validate() error {
	return dara.Validate(s)
}
