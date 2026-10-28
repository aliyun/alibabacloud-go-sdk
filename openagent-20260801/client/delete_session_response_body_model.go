// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteSessionResponseBodyData) *DeleteSessionResponseBody
	GetData() *DeleteSessionResponseBodyData
	SetRequestId(v string) *DeleteSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSessionResponseBody
	GetSuccess() *bool
}

type DeleteSessionResponseBody struct {
	// Returns an empty object {} upon successful deletion.
	//
	// example:
	//
	// {}
	Data *DeleteSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSessionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSessionResponseBody) GetData() *DeleteSessionResponseBodyData {
	return s.Data
}

func (s *DeleteSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSessionResponseBody) SetData(v *DeleteSessionResponseBodyData) *DeleteSessionResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSessionResponseBody) SetRequestId(v string) *DeleteSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSessionResponseBody) SetSuccess(v bool) *DeleteSessionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSessionResponseBodyData struct {
}

func (s DeleteSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSessionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
