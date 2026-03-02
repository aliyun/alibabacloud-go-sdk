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
	AccessKeyLastUsed *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed `json:"AccessKeyLastUsed,omitempty" xml:"AccessKeyLastUsed,omitempty" type:"Struct"`
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	LastUsedDate *string `json:"LastUsedDate,omitempty" xml:"LastUsedDate,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
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
