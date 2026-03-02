// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyInfoInRecycleBinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKey(v *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) *GetAccessKeyInfoInRecycleBinResponseBody
	GetAccessKey() *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey
	SetRequestId(v string) *GetAccessKeyInfoInRecycleBinResponseBody
	GetRequestId() *string
}

type GetAccessKeyInfoInRecycleBinResponseBody struct {
	AccessKey *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyInfoInRecycleBinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyInfoInRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyInfoInRecycleBinResponseBody) GetAccessKey() *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	return s.AccessKey
}

func (s *GetAccessKeyInfoInRecycleBinResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessKeyInfoInRecycleBinResponseBody) SetAccessKey(v *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) *GetAccessKeyInfoInRecycleBinResponseBody {
	s.AccessKey = v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBody) SetRequestId(v string) *GetAccessKeyInfoInRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBody) Validate() error {
	if s.AccessKey != nil {
		if err := s.AccessKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccessKeyInfoInRecycleBinResponseBodyAccessKey struct {
	AccessKeyId       *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	CreateDate        *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DeleteDate        *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	RecycleDate       *string `json:"RecycleDate,omitempty" xml:"RecycleDate,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
	UserRecycled      *bool   `json:"UserRecycled,omitempty" xml:"UserRecycled,omitempty"`
}

func (s GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) GoString() string {
	return s.String()
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) GetDeleteDate() *string {
	return s.DeleteDate
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) GetRecycleDate() *string {
	return s.RecycleDate
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) GetUserId() *string {
	return s.UserId
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) GetUserRecycled() *bool {
	return s.UserRecycled
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetAccessKeyId(v string) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.AccessKeyId = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetCreateDate(v string) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.CreateDate = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetDeleteDate(v string) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.DeleteDate = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetRecycleDate(v string) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.RecycleDate = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetUserId(v string) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.UserId = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetUserPrincipalName(v string) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.UserPrincipalName = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetUserRecycled(v bool) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.UserRecycled = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) Validate() error {
	return dara.Validate(s)
}
