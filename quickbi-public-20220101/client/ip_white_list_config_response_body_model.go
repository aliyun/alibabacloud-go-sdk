// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIpWhiteListConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *IpWhiteListConfigResponseBody
	GetRequestId() *string
	SetResult(v *IpWhiteListConfigResponseBodyResult) *IpWhiteListConfigResponseBody
	GetResult() *IpWhiteListConfigResponseBodyResult
	SetSuccess(v bool) *IpWhiteListConfigResponseBody
	GetSuccess() *bool
}

type IpWhiteListConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IP address whitelist.
	Result *IpWhiteListConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IpWhiteListConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IpWhiteListConfigResponseBody) GoString() string {
	return s.String()
}

func (s *IpWhiteListConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IpWhiteListConfigResponseBody) GetResult() *IpWhiteListConfigResponseBodyResult {
	return s.Result
}

func (s *IpWhiteListConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IpWhiteListConfigResponseBody) SetRequestId(v string) *IpWhiteListConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *IpWhiteListConfigResponseBody) SetResult(v *IpWhiteListConfigResponseBodyResult) *IpWhiteListConfigResponseBody {
	s.Result = v
	return s
}

func (s *IpWhiteListConfigResponseBody) SetSuccess(v bool) *IpWhiteListConfigResponseBody {
	s.Success = &v
	return s
}

func (s *IpWhiteListConfigResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IpWhiteListConfigResponseBodyResult struct {
	// The IP address whitelist array.
	IpWhiteList []*string `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty" type:"Repeated"`
}

func (s IpWhiteListConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s IpWhiteListConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IpWhiteListConfigResponseBodyResult) GetIpWhiteList() []*string {
	return s.IpWhiteList
}

func (s *IpWhiteListConfigResponseBodyResult) SetIpWhiteList(v []*string) *IpWhiteListConfigResponseBodyResult {
	s.IpWhiteList = v
	return s
}

func (s *IpWhiteListConfigResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
