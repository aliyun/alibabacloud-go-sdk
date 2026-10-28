// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOncallContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTimestamp(v string) *ListOncallContactsRequest
	GetTimestamp() *string
}

type ListOncallContactsRequest struct {
	// The timestamp of the on-call shift to query. Unit: milliseconds.
	//
	// example:
	//
	// 1730183838000
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s ListOncallContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOncallContactsRequest) GoString() string {
	return s.String()
}

func (s *ListOncallContactsRequest) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListOncallContactsRequest) SetTimestamp(v string) *ListOncallContactsRequest {
	s.Timestamp = &v
	return s
}

func (s *ListOncallContactsRequest) Validate() error {
	return dara.Validate(s)
}
