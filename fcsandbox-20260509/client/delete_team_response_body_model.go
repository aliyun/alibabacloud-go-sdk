// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTeamResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTeamResponseBody
	GetRequestId() *string
}

type DeleteTeamResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTeamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTeamResponseBody) SetCode(v string) *DeleteTeamResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTeamResponseBody) SetMessage(v string) *DeleteTeamResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTeamResponseBody) SetRequestId(v string) *DeleteTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTeamResponseBody) Validate() error {
	return dara.Validate(s)
}
