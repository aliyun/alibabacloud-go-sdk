// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserHonorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHonors(v []*QueryUserHonorsResponseBodyHonors) *QueryUserHonorsResponseBody
	GetHonors() []*QueryUserHonorsResponseBodyHonors
	SetNextToken(v string) *QueryUserHonorsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryUserHonorsResponseBody
	GetRequestId() *string
}

type QueryUserHonorsResponseBody struct {
	Honors []*QueryUserHonorsResponseBodyHonors `json:"honors,omitempty" xml:"honors,omitempty" type:"Repeated"`
	// example:
	//
	// http-trigger-nodejs10.luoni-old.1431999136518149.cn-hangzhou.fc.devsapp.net
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryUserHonorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserHonorsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBody) GetHonors() []*QueryUserHonorsResponseBodyHonors {
	return s.Honors
}

func (s *QueryUserHonorsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryUserHonorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserHonorsResponseBody) SetHonors(v []*QueryUserHonorsResponseBodyHonors) *QueryUserHonorsResponseBody {
	s.Honors = v
	return s
}

func (s *QueryUserHonorsResponseBody) SetNextToken(v string) *QueryUserHonorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryUserHonorsResponseBody) SetRequestId(v string) *QueryUserHonorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserHonorsResponseBody) Validate() error {
	if s.Honors != nil {
		for _, item := range s.Honors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryUserHonorsResponseBodyHonors struct {
	// example:
	//
	// null
	ExpirationTime *int64                                           `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	GrantHistory   []*QueryUserHonorsResponseBodyHonorsGrantHistory `json:"grantHistory,omitempty" xml:"grantHistory,omitempty" type:"Repeated"`
	HonorDesc      *string                                          `json:"honorDesc,omitempty" xml:"honorDesc,omitempty"`
	// example:
	//
	// 21659398
	HonorId   *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
	HonorName *string `json:"honorName,omitempty" xml:"honorName,omitempty"`
}

func (s QueryUserHonorsResponseBodyHonors) String() string {
	return dara.Prettify(s)
}

func (s QueryUserHonorsResponseBodyHonors) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBodyHonors) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *QueryUserHonorsResponseBodyHonors) GetGrantHistory() []*QueryUserHonorsResponseBodyHonorsGrantHistory {
	return s.GrantHistory
}

func (s *QueryUserHonorsResponseBodyHonors) GetHonorDesc() *string {
	return s.HonorDesc
}

func (s *QueryUserHonorsResponseBodyHonors) GetHonorId() *string {
	return s.HonorId
}

func (s *QueryUserHonorsResponseBodyHonors) GetHonorName() *string {
	return s.HonorName
}

func (s *QueryUserHonorsResponseBodyHonors) SetExpirationTime(v int64) *QueryUserHonorsResponseBodyHonors {
	s.ExpirationTime = &v
	return s
}

func (s *QueryUserHonorsResponseBodyHonors) SetGrantHistory(v []*QueryUserHonorsResponseBodyHonorsGrantHistory) *QueryUserHonorsResponseBodyHonors {
	s.GrantHistory = v
	return s
}

func (s *QueryUserHonorsResponseBodyHonors) SetHonorDesc(v string) *QueryUserHonorsResponseBodyHonors {
	s.HonorDesc = &v
	return s
}

func (s *QueryUserHonorsResponseBodyHonors) SetHonorId(v string) *QueryUserHonorsResponseBodyHonors {
	s.HonorId = &v
	return s
}

func (s *QueryUserHonorsResponseBodyHonors) SetHonorName(v string) *QueryUserHonorsResponseBodyHonors {
	s.HonorName = &v
	return s
}

func (s *QueryUserHonorsResponseBodyHonors) Validate() error {
	if s.GrantHistory != nil {
		for _, item := range s.GrantHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryUserHonorsResponseBodyHonorsGrantHistory struct {
	// example:
	//
	// 12312312312312312
	GrantTime *int64 `json:"grantTime,omitempty" xml:"grantTime,omitempty"`
	// example:
	//
	// 363784
	SenderUserid *string `json:"senderUserid,omitempty" xml:"senderUserid,omitempty"`
}

func (s QueryUserHonorsResponseBodyHonorsGrantHistory) String() string {
	return dara.Prettify(s)
}

func (s QueryUserHonorsResponseBodyHonorsGrantHistory) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBodyHonorsGrantHistory) GetGrantTime() *int64 {
	return s.GrantTime
}

func (s *QueryUserHonorsResponseBodyHonorsGrantHistory) GetSenderUserid() *string {
	return s.SenderUserid
}

func (s *QueryUserHonorsResponseBodyHonorsGrantHistory) SetGrantTime(v int64) *QueryUserHonorsResponseBodyHonorsGrantHistory {
	s.GrantTime = &v
	return s
}

func (s *QueryUserHonorsResponseBodyHonorsGrantHistory) SetSenderUserid(v string) *QueryUserHonorsResponseBodyHonorsGrantHistory {
	s.SenderUserid = &v
	return s
}

func (s *QueryUserHonorsResponseBodyHonorsGrantHistory) Validate() error {
	return dara.Validate(s)
}
