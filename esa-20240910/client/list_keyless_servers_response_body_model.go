// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeylessServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListKeylessServersResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListKeylessServersResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListKeylessServersResponseBody
	GetRequestId() *string
	SetResult(v []*ListKeylessServersResponseBodyResult) *ListKeylessServersResponseBody
	GetResult() []*ListKeylessServersResponseBodyResult
	SetSiteId(v int64) *ListKeylessServersResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *ListKeylessServersResponseBody
	GetSiteName() *string
	SetTotalCount(v int64) *ListKeylessServersResponseBody
	GetTotalCount() *int64
}

type ListKeylessServersResponseBody struct {
	PageNumber *int64                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*ListKeylessServersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	SiteId     *int64                                  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName   *string                                 `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKeylessServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKeylessServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeylessServersResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListKeylessServersResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListKeylessServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKeylessServersResponseBody) GetResult() []*ListKeylessServersResponseBodyResult {
	return s.Result
}

func (s *ListKeylessServersResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListKeylessServersResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *ListKeylessServersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListKeylessServersResponseBody) SetPageNumber(v int64) *ListKeylessServersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKeylessServersResponseBody) SetPageSize(v int64) *ListKeylessServersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKeylessServersResponseBody) SetRequestId(v string) *ListKeylessServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKeylessServersResponseBody) SetResult(v []*ListKeylessServersResponseBodyResult) *ListKeylessServersResponseBody {
	s.Result = v
	return s
}

func (s *ListKeylessServersResponseBody) SetSiteId(v int64) *ListKeylessServersResponseBody {
	s.SiteId = &v
	return s
}

func (s *ListKeylessServersResponseBody) SetSiteName(v string) *ListKeylessServersResponseBody {
	s.SiteName = &v
	return s
}

func (s *ListKeylessServersResponseBody) SetTotalCount(v int64) *ListKeylessServersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKeylessServersResponseBody) Validate() error {
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

type ListKeylessServersResponseBodyResult struct {
	CaCertificate     *string `json:"CaCertificate,omitempty" xml:"CaCertificate,omitempty"`
	ClientCertificate *string `json:"ClientCertificate,omitempty" xml:"ClientCertificate,omitempty"`
	ClientPrivateKey  *string `json:"ClientPrivateKey,omitempty" xml:"ClientPrivateKey,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Host              *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Id                *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Port              *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	UpdateTime        *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Verify            *bool   `json:"Verify,omitempty" xml:"Verify,omitempty"`
}

func (s ListKeylessServersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListKeylessServersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListKeylessServersResponseBodyResult) GetCaCertificate() *string {
	return s.CaCertificate
}

func (s *ListKeylessServersResponseBodyResult) GetClientCertificate() *string {
	return s.ClientCertificate
}

func (s *ListKeylessServersResponseBodyResult) GetClientPrivateKey() *string {
	return s.ClientPrivateKey
}

func (s *ListKeylessServersResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListKeylessServersResponseBodyResult) GetHost() *string {
	return s.Host
}

func (s *ListKeylessServersResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListKeylessServersResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListKeylessServersResponseBodyResult) GetPort() *int64 {
	return s.Port
}

func (s *ListKeylessServersResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListKeylessServersResponseBodyResult) GetVerify() *bool {
	return s.Verify
}

func (s *ListKeylessServersResponseBodyResult) SetCaCertificate(v string) *ListKeylessServersResponseBodyResult {
	s.CaCertificate = &v
	return s
}

func (s *ListKeylessServersResponseBodyResult) SetClientCertificate(v string) *ListKeylessServersResponseBodyResult {
	s.ClientCertificate = &v
	return s
}

func (s *ListKeylessServersResponseBodyResult) SetClientPrivateKey(v string) *ListKeylessServersResponseBodyResult {
	s.ClientPrivateKey = &v
	return s
}

func (s *ListKeylessServersResponseBodyResult) SetCreateTime(v string) *ListKeylessServersResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListKeylessServersResponseBodyResult) SetHost(v string) *ListKeylessServersResponseBodyResult {
	s.Host = &v
	return s
}

func (s *ListKeylessServersResponseBodyResult) SetId(v string) *ListKeylessServersResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListKeylessServersResponseBodyResult) SetName(v string) *ListKeylessServersResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListKeylessServersResponseBodyResult) SetPort(v int64) *ListKeylessServersResponseBodyResult {
	s.Port = &v
	return s
}

func (s *ListKeylessServersResponseBodyResult) SetUpdateTime(v string) *ListKeylessServersResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *ListKeylessServersResponseBodyResult) SetVerify(v bool) *ListKeylessServersResponseBodyResult {
	s.Verify = &v
	return s
}

func (s *ListKeylessServersResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
