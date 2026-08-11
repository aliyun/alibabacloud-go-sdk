// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditContent(v string) *CreateAuditRequest
	GetAuditContent() *string
}

type CreateAuditRequest struct {
	// The array of review content. A maximum of 20 videos can be reviewed at a time. Convert the array to a string before passing it as the parameter value. For the specific parameter structure, see the AuditContent table below.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "MediaId": "93ab850b4f*****b54b6e91d24d81d4",
	//
	//             "Status": "Normal"
	//
	//       },
	//
	//       {
	//
	//             "MediaId": "f867fbfb58*****8bbab65c4480ae1d",
	//
	//             "Status": "Blocked",
	//
	//             "Reason": "xxxx",
	//
	//             "Comment": "xxxx"
	//
	//       }
	//
	// ]
	AuditContent *string `json:"AuditContent,omitempty" xml:"AuditContent,omitempty"`
}

func (s CreateAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAuditRequest) GoString() string {
	return s.String()
}

func (s *CreateAuditRequest) GetAuditContent() *string {
	return s.AuditContent
}

func (s *CreateAuditRequest) SetAuditContent(v string) *CreateAuditRequest {
	s.AuditContent = &v
	return s
}

func (s *CreateAuditRequest) Validate() error {
	return dara.Validate(s)
}
