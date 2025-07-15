// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventSubRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListEventSubRequest
	GetAppId() *string
}

type ListEventSubRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9qb1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListEventSubRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventSubRequest) GoString() string {
	return s.String()
}

func (s *ListEventSubRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListEventSubRequest) SetAppId(v string) *ListEventSubRequest {
	s.AppId = &v
	return s
}

func (s *ListEventSubRequest) Validate() error {
	return dara.Validate(s)
}
