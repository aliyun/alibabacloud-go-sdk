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
	// 5CCE804C-6450-49A7-B1DB-2460F7A97416
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
	// 2020-10-21T06:37:40Z
	LastUsedDate *string `json:"LastUsedDate,omitempty" xml:"LastUsedDate,omitempty"`
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

func (s *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) SetLastUsedDate(v string) *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed {
	s.LastUsedDate = &v
	return s
}

func (s *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) Validate() error {
	return dara.Validate(s)
}
