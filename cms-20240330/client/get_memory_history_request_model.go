// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMemoryHistoryRequest struct {
}

func (s GetMemoryHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryHistoryRequest) Validate() error {
	return dara.Validate(s)
}
