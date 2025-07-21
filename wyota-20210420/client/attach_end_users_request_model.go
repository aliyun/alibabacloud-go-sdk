// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEndUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserIds(v string) *AttachEndUsersRequest
	GetEndUserIds() *string
	SetSerialNo(v string) *AttachEndUsersRequest
	GetSerialNo() *string
}

type AttachEndUsersRequest struct {
	// This parameter is required.
	EndUserIds *string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s AttachEndUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachEndUsersRequest) GoString() string {
	return s.String()
}

func (s *AttachEndUsersRequest) GetEndUserIds() *string {
	return s.EndUserIds
}

func (s *AttachEndUsersRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *AttachEndUsersRequest) SetEndUserIds(v string) *AttachEndUsersRequest {
	s.EndUserIds = &v
	return s
}

func (s *AttachEndUsersRequest) SetSerialNo(v string) *AttachEndUsersRequest {
	s.SerialNo = &v
	return s
}

func (s *AttachEndUsersRequest) Validate() error {
	return dara.Validate(s)
}
