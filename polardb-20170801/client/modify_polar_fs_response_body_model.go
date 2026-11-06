// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarFsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPolarFsResponseBody
	GetRequestId() *string
}

type ModifyPolarFsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CED079B7-A408-41A1-BFF1-EC608E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolarFsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarFsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolarFsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPolarFsResponseBody) SetRequestId(v string) *ModifyPolarFsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPolarFsResponseBody) Validate() error {
	return dara.Validate(s)
}
