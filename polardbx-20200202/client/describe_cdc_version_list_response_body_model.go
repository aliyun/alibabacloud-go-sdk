// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdcVersionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeCdcVersionListResponseBodyData) *DescribeCdcVersionListResponseBody
	GetData() *DescribeCdcVersionListResponseBodyData
	SetRequestId(v string) *DescribeCdcVersionListResponseBody
	GetRequestId() *string
}

type DescribeCdcVersionListResponseBody struct {
	Data *DescribeCdcVersionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 97632117-E477-5FE8-B424-F059CBBDD919
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdcVersionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcVersionListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdcVersionListResponseBody) GetData() *DescribeCdcVersionListResponseBodyData {
	return s.Data
}

func (s *DescribeCdcVersionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdcVersionListResponseBody) SetData(v *DescribeCdcVersionListResponseBodyData) *DescribeCdcVersionListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCdcVersionListResponseBody) SetRequestId(v string) *DescribeCdcVersionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdcVersionListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdcVersionListResponseBodyData struct {
	VersionList []*string `json:"VersionList,omitempty" xml:"VersionList,omitempty" type:"Repeated"`
}

func (s DescribeCdcVersionListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcVersionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCdcVersionListResponseBodyData) GetVersionList() []*string {
	return s.VersionList
}

func (s *DescribeCdcVersionListResponseBodyData) SetVersionList(v []*string) *DescribeCdcVersionListResponseBodyData {
	s.VersionList = v
	return s
}

func (s *DescribeCdcVersionListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
