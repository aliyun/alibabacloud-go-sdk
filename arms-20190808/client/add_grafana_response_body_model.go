// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGrafanaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *AddGrafanaResponseBody
	GetData() *string
	SetRequestId(v string) *AddGrafanaResponseBody
	GetRequestId() *string
}

type AddGrafanaResponseBody struct {
	// Indicates whether the call was successful.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddGrafanaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGrafanaResponseBody) GoString() string {
	return s.String()
}

func (s *AddGrafanaResponseBody) GetData() *string {
	return s.Data
}

func (s *AddGrafanaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGrafanaResponseBody) SetData(v string) *AddGrafanaResponseBody {
	s.Data = &v
	return s
}

func (s *AddGrafanaResponseBody) SetRequestId(v string) *AddGrafanaResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGrafanaResponseBody) Validate() error {
	return dara.Validate(s)
}
