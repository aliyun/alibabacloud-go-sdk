// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryId(v string) *SubmitQueryResponseBody
	GetQueryId() *string
}

type SubmitQueryResponseBody struct {
	QueryId *string `json:"queryId,omitempty" xml:"queryId,omitempty"`
}

func (s SubmitQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitQueryResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitQueryResponseBody) GetQueryId() *string {
	return s.QueryId
}

func (s *SubmitQueryResponseBody) SetQueryId(v string) *SubmitQueryResponseBody {
	s.QueryId = &v
	return s
}

func (s *SubmitQueryResponseBody) Validate() error {
	return dara.Validate(s)
}
