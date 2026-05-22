// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyCertificateResponseBody
	GetRequestId() *string
	SetResult(v []*ApplyCertificateResponseBodyResult) *ApplyCertificateResponseBody
	GetResult() []*ApplyCertificateResponseBodyResult
	SetSiteName(v string) *ApplyCertificateResponseBody
	GetSiteName() *string
	SetTotalCount(v int64) *ApplyCertificateResponseBody
	GetTotalCount() *int64
}

type ApplyCertificateResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*ApplyCertificateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	SiteName   *string                               `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ApplyCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyCertificateResponseBody) GetResult() []*ApplyCertificateResponseBodyResult {
	return s.Result
}

func (s *ApplyCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *ApplyCertificateResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ApplyCertificateResponseBody) SetRequestId(v string) *ApplyCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyCertificateResponseBody) SetResult(v []*ApplyCertificateResponseBodyResult) *ApplyCertificateResponseBody {
	s.Result = v
	return s
}

func (s *ApplyCertificateResponseBody) SetSiteName(v string) *ApplyCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *ApplyCertificateResponseBody) SetTotalCount(v int64) *ApplyCertificateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ApplyCertificateResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyCertificateResponseBodyResult struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ApplyCertificateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ApplyCertificateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ApplyCertificateResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *ApplyCertificateResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ApplyCertificateResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ApplyCertificateResponseBodyResult) SetDomain(v string) *ApplyCertificateResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ApplyCertificateResponseBodyResult) SetId(v string) *ApplyCertificateResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ApplyCertificateResponseBodyResult) SetStatus(v string) *ApplyCertificateResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ApplyCertificateResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
