// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceLogsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetWebLogEntrys(v []*WebLogEntry) *DescribeInstanceLogsOutput
	GetWebLogEntrys() []*WebLogEntry
}

type DescribeInstanceLogsOutput struct {
	WebLogEntrys []*WebLogEntry `json:"WebLogEntrys,omitempty" xml:"WebLogEntrys,omitempty" type:"Repeated"`
}

func (s DescribeInstanceLogsOutput) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLogsOutput) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLogsOutput) GetWebLogEntrys() []*WebLogEntry {
	return s.WebLogEntrys
}

func (s *DescribeInstanceLogsOutput) SetWebLogEntrys(v []*WebLogEntry) *DescribeInstanceLogsOutput {
	s.WebLogEntrys = v
	return s
}

func (s *DescribeInstanceLogsOutput) Validate() error {
	if s.WebLogEntrys != nil {
		for _, item := range s.WebLogEntrys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
