// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrinterEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*DescribePrinterEventsResponseBodyEvents) *DescribePrinterEventsResponseBody
	GetEvents() []*DescribePrinterEventsResponseBodyEvents
	SetNextToken(v string) *DescribePrinterEventsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribePrinterEventsResponseBody
	GetRequestId() *string
}

type DescribePrinterEventsResponseBody struct {
	// The user events.
	Events []*DescribePrinterEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The pagination token for the next query. If NextToken is empty, no more results exist.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E54EB497-D7B7-5F04-B744-D8DFA7B******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePrinterEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrinterEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrinterEventsResponseBody) GetEvents() []*DescribePrinterEventsResponseBodyEvents {
	return s.Events
}

func (s *DescribePrinterEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePrinterEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrinterEventsResponseBody) SetEvents(v []*DescribePrinterEventsResponseBodyEvents) *DescribePrinterEventsResponseBody {
	s.Events = v
	return s
}

func (s *DescribePrinterEventsResponseBody) SetNextToken(v string) *DescribePrinterEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePrinterEventsResponseBody) SetRequestId(v string) *DescribePrinterEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrinterEventsResponseBody) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePrinterEventsResponseBodyEvents struct {
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-8fupvkhg0aayu****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The cloud computer name.
	//
	// example:
	//
	// desktop-001
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The end user ID.
	//
	// example:
	//
	// user001
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 5651188b-3070-d1cc-5311-75753d59****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The printer driver name.
	//
	// example:
	//
	// HP LaserJet PCL 6
	PrinterDriver *string `json:"PrinterDriver,omitempty" xml:"PrinterDriver,omitempty"`
	// The number of copies to print.
	//
	// example:
	//
	// 1
	PrinterJobCopies *int32 `json:"PrinterJobCopies,omitempty" xml:"PrinterJobCopies,omitempty"`
	// The print job name.
	//
	// example:
	//
	// report.pdf
	PrinterJobName *string `json:"PrinterJobName,omitempty" xml:"PrinterJobName,omitempty"`
	// The total number of pages in the print job.
	//
	// example:
	//
	// 1
	PrinterJobPages *int32 `json:"PrinterJobPages,omitempty" xml:"PrinterJobPages,omitempty"`
	// The number of printed pages.
	//
	// example:
	//
	// 1
	PrinterJobPrintedPages *int32 `json:"PrinterJobPrintedPages,omitempty" xml:"PrinterJobPrintedPages,omitempty"`
	// The print job size, in bytes.
	//
	// example:
	//
	// 2632446
	PrinterJobSize *int64 `json:"PrinterJobSize,omitempty" xml:"PrinterJobSize,omitempty"`
	// The print job time, in millisecond-precision UNIX timestamp.
	//
	// example:
	//
	// 1706140800000
	PrinterJobTime *int64 `json:"PrinterJobTime,omitempty" xml:"PrinterJobTime,omitempty"`
	// The printer name.
	//
	// example:
	//
	// HP LaserJet Pro
	PrinterName *string `json:"PrinterName,omitempty" xml:"PrinterName,omitempty"`
	// The printer port.
	//
	// example:
	//
	// USB001
	PrinterPort *string `json:"PrinterPort,omitempty" xml:"PrinterPort,omitempty"`
	// The printer redirection type.
	//
	// example:
	//
	// 1
	PrinterRedirType *int32 `json:"PrinterRedirType,omitempty" xml:"PrinterRedirType,omitempty"`
}

func (s DescribePrinterEventsResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribePrinterEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribePrinterEventsResponseBodyEvents) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribePrinterEventsResponseBodyEvents) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribePrinterEventsResponseBodyEvents) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribePrinterEventsResponseBodyEvents) GetEventId() *string {
	return s.EventId
}

func (s *DescribePrinterEventsResponseBodyEvents) GetPrinterDriver() *string {
	return s.PrinterDriver
}

func (s *DescribePrinterEventsResponseBodyEvents) GetPrinterJobCopies() *int32 {
	return s.PrinterJobCopies
}

func (s *DescribePrinterEventsResponseBodyEvents) GetPrinterJobName() *string {
	return s.PrinterJobName
}

func (s *DescribePrinterEventsResponseBodyEvents) GetPrinterJobPages() *int32 {
	return s.PrinterJobPages
}

func (s *DescribePrinterEventsResponseBodyEvents) GetPrinterJobPrintedPages() *int32 {
	return s.PrinterJobPrintedPages
}

func (s *DescribePrinterEventsResponseBodyEvents) GetPrinterJobSize() *int64 {
	return s.PrinterJobSize
}

func (s *DescribePrinterEventsResponseBodyEvents) GetPrinterJobTime() *int64 {
	return s.PrinterJobTime
}

func (s *DescribePrinterEventsResponseBodyEvents) GetPrinterName() *string {
	return s.PrinterName
}

func (s *DescribePrinterEventsResponseBodyEvents) GetPrinterPort() *string {
	return s.PrinterPort
}

func (s *DescribePrinterEventsResponseBodyEvents) GetPrinterRedirType() *int32 {
	return s.PrinterRedirType
}

func (s *DescribePrinterEventsResponseBodyEvents) SetDesktopId(v string) *DescribePrinterEventsResponseBodyEvents {
	s.DesktopId = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetDesktopName(v string) *DescribePrinterEventsResponseBodyEvents {
	s.DesktopName = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetEndUserId(v string) *DescribePrinterEventsResponseBodyEvents {
	s.EndUserId = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetEventId(v string) *DescribePrinterEventsResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetPrinterDriver(v string) *DescribePrinterEventsResponseBodyEvents {
	s.PrinterDriver = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetPrinterJobCopies(v int32) *DescribePrinterEventsResponseBodyEvents {
	s.PrinterJobCopies = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetPrinterJobName(v string) *DescribePrinterEventsResponseBodyEvents {
	s.PrinterJobName = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetPrinterJobPages(v int32) *DescribePrinterEventsResponseBodyEvents {
	s.PrinterJobPages = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetPrinterJobPrintedPages(v int32) *DescribePrinterEventsResponseBodyEvents {
	s.PrinterJobPrintedPages = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetPrinterJobSize(v int64) *DescribePrinterEventsResponseBodyEvents {
	s.PrinterJobSize = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetPrinterJobTime(v int64) *DescribePrinterEventsResponseBodyEvents {
	s.PrinterJobTime = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetPrinterName(v string) *DescribePrinterEventsResponseBodyEvents {
	s.PrinterName = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetPrinterPort(v string) *DescribePrinterEventsResponseBodyEvents {
	s.PrinterPort = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) SetPrinterRedirType(v int32) *DescribePrinterEventsResponseBodyEvents {
	s.PrinterRedirType = &v
	return s
}

func (s *DescribePrinterEventsResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}
