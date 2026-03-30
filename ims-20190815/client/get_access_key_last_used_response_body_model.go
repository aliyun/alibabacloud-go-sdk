// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyLastUsed(v *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) *GetAccessKeyLastUsedResponseBody
	GetAccessKeyLastUsed() *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed
	SetRequestId(v string) *GetAccessKeyLastUsedResponseBody
	GetRequestId() *string
}

type GetAccessKeyLastUsedResponseBody struct {
	// The details of the time when the AccessKey pair was used for the last time.
	AccessKeyLastUsed *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed `json:"AccessKeyLastUsed,omitempty" xml:"AccessKeyLastUsed,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B29C79F6-354B-4297-A994-1338CC22A2EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyLastUsedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResponseBody) GetAccessKeyLastUsed() *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed {
	return s.AccessKeyLastUsed
}

func (s *GetAccessKeyLastUsedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessKeyLastUsedResponseBody) SetAccessKeyLastUsed(v *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) *GetAccessKeyLastUsedResponseBody {
	s.AccessKeyLastUsed = v
	return s
}

func (s *GetAccessKeyLastUsedResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessKeyLastUsedResponseBody) Validate() error {
	if s.AccessKeyLastUsed != nil {
		if err := s.AccessKeyLastUsed.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed struct {
	// The time when the AccessKey pair was used for the last time.
	//
	// example:
	//
	// 2020-10-16T01:37:37Z
	LastUsedDate *string `json:"LastUsedDate,omitempty" xml:"LastUsedDate,omitempty"`
	// The Alibaba Cloud service that was last accessed by using the AccessKey pair.
	//
	// example:
	//
	// Ram
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) GetLastUsedDate() *string {
	return s.LastUsedDate
}

func (s *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) SetLastUsedDate(v string) *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed {
	s.LastUsedDate = &v
	return s
}

func (s *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) SetServiceName(v string) *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed {
	s.ServiceName = &v
	return s
}

func (s *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) Validate() error {
	return dara.Validate(s)
}
