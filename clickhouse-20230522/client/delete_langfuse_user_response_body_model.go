// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteLangfuseUserResponseBodyData) *DeleteLangfuseUserResponseBody
	GetData() *DeleteLangfuseUserResponseBodyData
	SetRequestId(v string) *DeleteLangfuseUserResponseBody
	GetRequestId() *string
}

type DeleteLangfuseUserResponseBody struct {
	// The response result.
	Data *DeleteLangfuseUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLangfuseUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseUserResponseBody) GetData() *DeleteLangfuseUserResponseBodyData {
	return s.Data
}

func (s *DeleteLangfuseUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLangfuseUserResponseBody) SetData(v *DeleteLangfuseUserResponseBodyData) *DeleteLangfuseUserResponseBody {
	s.Data = v
	return s
}

func (s *DeleteLangfuseUserResponseBody) SetRequestId(v string) *DeleteLangfuseUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLangfuseUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteLangfuseUserResponseBodyData struct {
	// The email address of the user.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s DeleteLangfuseUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseUserResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *DeleteLangfuseUserResponseBodyData) SetEmail(v string) *DeleteLangfuseUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *DeleteLangfuseUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
