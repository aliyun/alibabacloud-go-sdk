// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMosQueryRegistrationInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v string) *MosQueryRegistrationInfoRequest
	GetActivityId() *string
	SetCheckinStatus(v string) *MosQueryRegistrationInfoRequest
	GetCheckinStatus() *string
	SetKeyword(v string) *MosQueryRegistrationInfoRequest
	GetKeyword() *string
	SetPage(v string) *MosQueryRegistrationInfoRequest
	GetPage() *string
	SetPageSize(v string) *MosQueryRegistrationInfoRequest
	GetPageSize() *string
	SetQrCode(v string) *MosQueryRegistrationInfoRequest
	GetQrCode() *string
}

type MosQueryRegistrationInfoRequest struct {
	ActivityId    *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	CheckinStatus *string `json:"CheckinStatus,omitempty" xml:"CheckinStatus,omitempty"`
	Keyword       *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Page          *string `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QrCode        *string `json:"QrCode,omitempty" xml:"QrCode,omitempty"`
}

func (s MosQueryRegistrationInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s MosQueryRegistrationInfoRequest) GoString() string {
	return s.String()
}

func (s *MosQueryRegistrationInfoRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *MosQueryRegistrationInfoRequest) GetCheckinStatus() *string {
	return s.CheckinStatus
}

func (s *MosQueryRegistrationInfoRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *MosQueryRegistrationInfoRequest) GetPage() *string {
	return s.Page
}

func (s *MosQueryRegistrationInfoRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *MosQueryRegistrationInfoRequest) GetQrCode() *string {
	return s.QrCode
}

func (s *MosQueryRegistrationInfoRequest) SetActivityId(v string) *MosQueryRegistrationInfoRequest {
	s.ActivityId = &v
	return s
}

func (s *MosQueryRegistrationInfoRequest) SetCheckinStatus(v string) *MosQueryRegistrationInfoRequest {
	s.CheckinStatus = &v
	return s
}

func (s *MosQueryRegistrationInfoRequest) SetKeyword(v string) *MosQueryRegistrationInfoRequest {
	s.Keyword = &v
	return s
}

func (s *MosQueryRegistrationInfoRequest) SetPage(v string) *MosQueryRegistrationInfoRequest {
	s.Page = &v
	return s
}

func (s *MosQueryRegistrationInfoRequest) SetPageSize(v string) *MosQueryRegistrationInfoRequest {
	s.PageSize = &v
	return s
}

func (s *MosQueryRegistrationInfoRequest) SetQrCode(v string) *MosQueryRegistrationInfoRequest {
	s.QrCode = &v
	return s
}

func (s *MosQueryRegistrationInfoRequest) Validate() error {
	return dara.Validate(s)
}
