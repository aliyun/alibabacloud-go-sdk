// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEndUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserIds(v string) *DetachEndUsersRequest
	GetEndUserIds() *string
	SetSerialNo(v string) *DetachEndUsersRequest
	GetSerialNo() *string
}

type DetachEndUsersRequest struct {
	// This parameter is required.
	EndUserIds *string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s DetachEndUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachEndUsersRequest) GoString() string {
	return s.String()
}

func (s *DetachEndUsersRequest) GetEndUserIds() *string {
	return s.EndUserIds
}

func (s *DetachEndUsersRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *DetachEndUsersRequest) SetEndUserIds(v string) *DetachEndUsersRequest {
	s.EndUserIds = &v
	return s
}

func (s *DetachEndUsersRequest) SetSerialNo(v string) *DetachEndUsersRequest {
	s.SerialNo = &v
	return s
}

func (s *DetachEndUsersRequest) Validate() error {
	return dara.Validate(s)
}
