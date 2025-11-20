// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgHonorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *QueryOrgHonorsResponseBody
	GetNextToken() *string
	SetOpenHonors(v []*QueryOrgHonorsResponseBodyOpenHonors) *QueryOrgHonorsResponseBody
	GetOpenHonors() []*QueryOrgHonorsResponseBodyOpenHonors
	SetRequestId(v string) *QueryOrgHonorsResponseBody
	GetRequestId() *string
}

type QueryOrgHonorsResponseBody struct {
	// example:
	//
	// 15
	NextToken  *string                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenHonors []*QueryOrgHonorsResponseBodyOpenHonors `json:"openHonors,omitempty" xml:"openHonors,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryOrgHonorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgHonorsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryOrgHonorsResponseBody) GetOpenHonors() []*QueryOrgHonorsResponseBodyOpenHonors {
	return s.OpenHonors
}

func (s *QueryOrgHonorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryOrgHonorsResponseBody) SetNextToken(v string) *QueryOrgHonorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryOrgHonorsResponseBody) SetOpenHonors(v []*QueryOrgHonorsResponseBodyOpenHonors) *QueryOrgHonorsResponseBody {
	s.OpenHonors = v
	return s
}

func (s *QueryOrgHonorsResponseBody) SetRequestId(v string) *QueryOrgHonorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrgHonorsResponseBody) Validate() error {
	if s.OpenHonors != nil {
		for _, item := range s.OpenHonors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryOrgHonorsResponseBodyOpenHonors struct {
	HonorDesc *string `json:"honorDesc,omitempty" xml:"honorDesc,omitempty"`
	// example:
	//
	// 21658579
	HonorId *int64 `json:"honorId,omitempty" xml:"honorId,omitempty"`
	// example:
	//
	// https://xxxx.ali-cdn.com/jfejfi.jpg
	HonorImgUrl *string `json:"honorImgUrl,omitempty" xml:"honorImgUrl,omitempty"`
	HonorName   *string `json:"honorName,omitempty" xml:"honorName,omitempty"`
	// example:
	//
	// https://xxxx.ali-cdn.com/jfejfi.jpg
	HonorPendantImgUrl *string `json:"honorPendantImgUrl,omitempty" xml:"honorPendantImgUrl,omitempty"`
}

func (s QueryOrgHonorsResponseBodyOpenHonors) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgHonorsResponseBodyOpenHonors) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponseBodyOpenHonors) GetHonorDesc() *string {
	return s.HonorDesc
}

func (s *QueryOrgHonorsResponseBodyOpenHonors) GetHonorId() *int64 {
	return s.HonorId
}

func (s *QueryOrgHonorsResponseBodyOpenHonors) GetHonorImgUrl() *string {
	return s.HonorImgUrl
}

func (s *QueryOrgHonorsResponseBodyOpenHonors) GetHonorName() *string {
	return s.HonorName
}

func (s *QueryOrgHonorsResponseBodyOpenHonors) GetHonorPendantImgUrl() *string {
	return s.HonorPendantImgUrl
}

func (s *QueryOrgHonorsResponseBodyOpenHonors) SetHonorDesc(v string) *QueryOrgHonorsResponseBodyOpenHonors {
	s.HonorDesc = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyOpenHonors) SetHonorId(v int64) *QueryOrgHonorsResponseBodyOpenHonors {
	s.HonorId = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyOpenHonors) SetHonorImgUrl(v string) *QueryOrgHonorsResponseBodyOpenHonors {
	s.HonorImgUrl = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyOpenHonors) SetHonorName(v string) *QueryOrgHonorsResponseBodyOpenHonors {
	s.HonorName = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyOpenHonors) SetHonorPendantImgUrl(v string) *QueryOrgHonorsResponseBodyOpenHonors {
	s.HonorPendantImgUrl = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyOpenHonors) Validate() error {
	return dara.Validate(s)
}
