// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackInstanceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *RollbackInstanceVersionResponseBody
	GetData() *string
	SetRequestId(v string) *RollbackInstanceVersionResponseBody
	GetRequestId() *string
}

type RollbackInstanceVersionResponseBody struct {
	// Indicates whether the instance version was rolled back.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DSSDF-SEWE-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackInstanceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackInstanceVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *RollbackInstanceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackInstanceVersionResponseBody) SetData(v string) *RollbackInstanceVersionResponseBody {
	s.Data = &v
	return s
}

func (s *RollbackInstanceVersionResponseBody) SetRequestId(v string) *RollbackInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackInstanceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
