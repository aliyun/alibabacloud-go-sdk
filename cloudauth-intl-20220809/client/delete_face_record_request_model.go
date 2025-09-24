// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteFaceRecordRequest
	GetId() *string
}

type DeleteFaceRecordRequest struct {
	// Primary Key ID
	//
	// example:
	//
	// 344537
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteFaceRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceRecordRequest) GetId() *string {
	return s.Id
}

func (s *DeleteFaceRecordRequest) SetId(v string) *DeleteFaceRecordRequest {
	s.Id = &v
	return s
}

func (s *DeleteFaceRecordRequest) Validate() error {
	return dara.Validate(s)
}
