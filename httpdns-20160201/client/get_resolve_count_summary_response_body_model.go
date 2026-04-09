// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResolveCountSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResolveCountSummaryResponseBody
	GetRequestId() *string
	SetResolveSummary(v *GetResolveCountSummaryResponseBodyResolveSummary) *GetResolveCountSummaryResponseBody
	GetResolveSummary() *GetResolveCountSummaryResponseBodyResolveSummary
}

type GetResolveCountSummaryResponseBody struct {
	// example:
	//
	// 3106FFF3-3612-542A-B2CD-3CF4CD48****
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResolveSummary *GetResolveCountSummaryResponseBodyResolveSummary `json:"ResolveSummary,omitempty" xml:"ResolveSummary,omitempty" type:"Struct"`
}

func (s GetResolveCountSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResolveCountSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetResolveCountSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResolveCountSummaryResponseBody) GetResolveSummary() *GetResolveCountSummaryResponseBodyResolveSummary {
	return s.ResolveSummary
}

func (s *GetResolveCountSummaryResponseBody) SetRequestId(v string) *GetResolveCountSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResolveCountSummaryResponseBody) SetResolveSummary(v *GetResolveCountSummaryResponseBodyResolveSummary) *GetResolveCountSummaryResponseBody {
	s.ResolveSummary = v
	return s
}

func (s *GetResolveCountSummaryResponseBody) Validate() error {
	if s.ResolveSummary != nil {
		if err := s.ResolveSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResolveCountSummaryResponseBodyResolveSummary struct {
	Doh *int64 `json:"Doh,omitempty" xml:"Doh,omitempty"`
	// example:
	//
	// 123
	Http *int64 `json:"Http,omitempty" xml:"Http,omitempty"`
	// example:
	//
	// 123
	Http6   *int64  `json:"Http6,omitempty" xml:"Http6,omitempty"`
	HttpAes *string `json:"HttpAes,omitempty" xml:"HttpAes,omitempty"`
	// example:
	//
	// 123
	Https *int64 `json:"Https,omitempty" xml:"Https,omitempty"`
	// example:
	//
	// 123
	Https6   *int64  `json:"Https6,omitempty" xml:"Https6,omitempty"`
	HttpsAes *string `json:"HttpsAes,omitempty" xml:"HttpsAes,omitempty"`
}

func (s GetResolveCountSummaryResponseBodyResolveSummary) String() string {
	return dara.Prettify(s)
}

func (s GetResolveCountSummaryResponseBodyResolveSummary) GoString() string {
	return s.String()
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) GetDoh() *int64 {
	return s.Doh
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) GetHttp() *int64 {
	return s.Http
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) GetHttp6() *int64 {
	return s.Http6
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) GetHttpAes() *string {
	return s.HttpAes
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) GetHttps() *int64 {
	return s.Https
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) GetHttps6() *int64 {
	return s.Https6
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) GetHttpsAes() *string {
	return s.HttpsAes
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) SetDoh(v int64) *GetResolveCountSummaryResponseBodyResolveSummary {
	s.Doh = &v
	return s
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) SetHttp(v int64) *GetResolveCountSummaryResponseBodyResolveSummary {
	s.Http = &v
	return s
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) SetHttp6(v int64) *GetResolveCountSummaryResponseBodyResolveSummary {
	s.Http6 = &v
	return s
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) SetHttpAes(v string) *GetResolveCountSummaryResponseBodyResolveSummary {
	s.HttpAes = &v
	return s
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) SetHttps(v int64) *GetResolveCountSummaryResponseBodyResolveSummary {
	s.Https = &v
	return s
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) SetHttps6(v int64) *GetResolveCountSummaryResponseBodyResolveSummary {
	s.Https6 = &v
	return s
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) SetHttpsAes(v string) *GetResolveCountSummaryResponseBodyResolveSummary {
	s.HttpsAes = &v
	return s
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) Validate() error {
	return dara.Validate(s)
}
