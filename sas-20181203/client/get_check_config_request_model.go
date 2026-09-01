// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceDirectoryAccountId(v int64) *GetCheckConfigRequest
	GetResourceDirectoryAccountId() *int64
}

type GetCheckConfigRequest struct {
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s GetCheckConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConfigRequest) GoString() string {
	return s.String()
}

func (s *GetCheckConfigRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *GetCheckConfigRequest) SetResourceDirectoryAccountId(v int64) *GetCheckConfigRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *GetCheckConfigRequest) Validate() error {
	return dara.Validate(s)
}
