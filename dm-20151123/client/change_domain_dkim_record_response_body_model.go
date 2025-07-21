// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDomainDkimRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChanged(v bool) *ChangeDomainDkimRecordResponseBody
	GetChanged() *bool
	SetDkimPublicKey(v string) *ChangeDomainDkimRecordResponseBody
	GetDkimPublicKey() *string
	SetDkimRsaLength(v int32) *ChangeDomainDkimRecordResponseBody
	GetDkimRsaLength() *int32
	SetHostname(v string) *ChangeDomainDkimRecordResponseBody
	GetHostname() *string
	SetRequestId(v string) *ChangeDomainDkimRecordResponseBody
	GetRequestId() *string
}

type ChangeDomainDkimRecordResponseBody struct {
	Changed       *bool   `json:"Changed,omitempty" xml:"Changed,omitempty"`
	DkimPublicKey *string `json:"DkimPublicKey,omitempty" xml:"DkimPublicKey,omitempty"`
	DkimRsaLength *int32  `json:"DkimRsaLength,omitempty" xml:"DkimRsaLength,omitempty"`
	Hostname      *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeDomainDkimRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeDomainDkimRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDomainDkimRecordResponseBody) GetChanged() *bool {
	return s.Changed
}

func (s *ChangeDomainDkimRecordResponseBody) GetDkimPublicKey() *string {
	return s.DkimPublicKey
}

func (s *ChangeDomainDkimRecordResponseBody) GetDkimRsaLength() *int32 {
	return s.DkimRsaLength
}

func (s *ChangeDomainDkimRecordResponseBody) GetHostname() *string {
	return s.Hostname
}

func (s *ChangeDomainDkimRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeDomainDkimRecordResponseBody) SetChanged(v bool) *ChangeDomainDkimRecordResponseBody {
	s.Changed = &v
	return s
}

func (s *ChangeDomainDkimRecordResponseBody) SetDkimPublicKey(v string) *ChangeDomainDkimRecordResponseBody {
	s.DkimPublicKey = &v
	return s
}

func (s *ChangeDomainDkimRecordResponseBody) SetDkimRsaLength(v int32) *ChangeDomainDkimRecordResponseBody {
	s.DkimRsaLength = &v
	return s
}

func (s *ChangeDomainDkimRecordResponseBody) SetHostname(v string) *ChangeDomainDkimRecordResponseBody {
	s.Hostname = &v
	return s
}

func (s *ChangeDomainDkimRecordResponseBody) SetRequestId(v string) *ChangeDomainDkimRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDomainDkimRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
