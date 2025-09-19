// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDiscoverDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateMark(v string) *QueryDiscoverDatabaseRequest
	GetCreateMark() *string
}

type QueryDiscoverDatabaseRequest struct {
	// The ID of the scan task.
	//
	// > You can call the [StartDiscoverDatabaseTask](~~StartDiscoverDatabaseTask~~) operation to query the ID of the task.
	//
	// example:
	//
	// 7f7b051f-7d1c-46da-b253-a03f3a27****
	CreateMark *string `json:"CreateMark,omitempty" xml:"CreateMark,omitempty"`
}

func (s QueryDiscoverDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDiscoverDatabaseRequest) GoString() string {
	return s.String()
}

func (s *QueryDiscoverDatabaseRequest) GetCreateMark() *string {
	return s.CreateMark
}

func (s *QueryDiscoverDatabaseRequest) SetCreateMark(v string) *QueryDiscoverDatabaseRequest {
	s.CreateMark = &v
	return s
}

func (s *QueryDiscoverDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
