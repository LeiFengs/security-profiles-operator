//go:build linux && !no_bpf
// +build linux,!no_bpf

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package recorderfakes

import (
	"io"
	"io/fs"
	"os"
	"os/exec"
	"sync"

	"github.com/aquasecurity/libbpfgo"
	seccompa "github.com/containers/common/pkg/seccomp"
	seccomp "github.com/seccomp/libseccomp-golang"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/printers"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/bpfrecorder"
)

type FakeImpl struct {
	CloseFileStub        func(*os.File)
	closeFileMutex       sync.RWMutex
	closeFileArgsForCall []struct {
		arg1 *os.File
	}
	CmdPidStub        func(*exec.Cmd) uint32
	cmdPidMutex       sync.RWMutex
	cmdPidArgsForCall []struct {
		arg1 *exec.Cmd
	}
	cmdPidReturns struct {
		result1 uint32
	}
	cmdPidReturnsOnCall map[int]struct {
		result1 uint32
	}
	CmdStartStub        func(*exec.Cmd) error
	cmdStartMutex       sync.RWMutex
	cmdStartArgsForCall []struct {
		arg1 *exec.Cmd
	}
	cmdStartReturns struct {
		result1 error
	}
	cmdStartReturnsOnCall map[int]struct {
		result1 error
	}
	CmdWaitStub        func(*exec.Cmd) error
	cmdWaitMutex       sync.RWMutex
	cmdWaitArgsForCall []struct {
		arg1 *exec.Cmd
	}
	cmdWaitReturns struct {
		result1 error
	}
	cmdWaitReturnsOnCall map[int]struct {
		result1 error
	}
	CommandStub        func(string, ...string) *exec.Cmd
	commandMutex       sync.RWMutex
	commandArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	commandReturns struct {
		result1 *exec.Cmd
	}
	commandReturnsOnCall map[int]struct {
		result1 *exec.Cmd
	}
	CreateStub        func(string) (*os.File, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
	}
	createReturns struct {
		result1 *os.File
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *os.File
		result2 error
	}
	GetNameStub        func(seccomp.ScmpSyscall) (string, error)
	getNameMutex       sync.RWMutex
	getNameArgsForCall []struct {
		arg1 seccomp.ScmpSyscall
	}
	getNameReturns struct {
		result1 string
		result2 error
	}
	getNameReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GoArchToSeccompArchStub        func(string) (seccompa.Arch, error)
	goArchToSeccompArchMutex       sync.RWMutex
	goArchToSeccompArchArgsForCall []struct {
		arg1 string
	}
	goArchToSeccompArchReturns struct {
		result1 seccompa.Arch
		result2 error
	}
	goArchToSeccompArchReturnsOnCall map[int]struct {
		result1 seccompa.Arch
		result2 error
	}
	IteratorKeyStub        func(*libbpfgo.BPFMapIterator) []byte
	iteratorKeyMutex       sync.RWMutex
	iteratorKeyArgsForCall []struct {
		arg1 *libbpfgo.BPFMapIterator
	}
	iteratorKeyReturns struct {
		result1 []byte
	}
	iteratorKeyReturnsOnCall map[int]struct {
		result1 []byte
	}
	IteratorNextStub        func(*libbpfgo.BPFMapIterator) bool
	iteratorNextMutex       sync.RWMutex
	iteratorNextArgsForCall []struct {
		arg1 *libbpfgo.BPFMapIterator
	}
	iteratorNextReturns struct {
		result1 bool
	}
	iteratorNextReturnsOnCall map[int]struct {
		result1 bool
	}
	LoadBpfRecorderStub        func(*bpfrecorder.BpfRecorder) error
	loadBpfRecorderMutex       sync.RWMutex
	loadBpfRecorderArgsForCall []struct {
		arg1 *bpfrecorder.BpfRecorder
	}
	loadBpfRecorderReturns struct {
		result1 error
	}
	loadBpfRecorderReturnsOnCall map[int]struct {
		result1 error
	}
	MarshalIndentStub        func(any, string, string) ([]byte, error)
	marshalIndentMutex       sync.RWMutex
	marshalIndentArgsForCall []struct {
		arg1 any
		arg2 string
		arg3 string
	}
	marshalIndentReturns struct {
		result1 []byte
		result2 error
	}
	marshalIndentReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	PrintObjStub        func(printers.YAMLPrinter, runtime.Object, io.Writer) error
	printObjMutex       sync.RWMutex
	printObjArgsForCall []struct {
		arg1 printers.YAMLPrinter
		arg2 runtime.Object
		arg3 io.Writer
	}
	printObjReturns struct {
		result1 error
	}
	printObjReturnsOnCall map[int]struct {
		result1 error
	}
	SyscallsGetValueStub        func(*bpfrecorder.BpfRecorder, uint32) ([]byte, error)
	syscallsGetValueMutex       sync.RWMutex
	syscallsGetValueArgsForCall []struct {
		arg1 *bpfrecorder.BpfRecorder
		arg2 uint32
	}
	syscallsGetValueReturns struct {
		result1 []byte
		result2 error
	}
	syscallsGetValueReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	SyscallsIteratorStub        func(*bpfrecorder.BpfRecorder) *libbpfgo.BPFMapIterator
	syscallsIteratorMutex       sync.RWMutex
	syscallsIteratorArgsForCall []struct {
		arg1 *bpfrecorder.BpfRecorder
	}
	syscallsIteratorReturns struct {
		result1 *libbpfgo.BPFMapIterator
	}
	syscallsIteratorReturnsOnCall map[int]struct {
		result1 *libbpfgo.BPFMapIterator
	}
	UnloadBpfRecorderStub        func(*bpfrecorder.BpfRecorder)
	unloadBpfRecorderMutex       sync.RWMutex
	unloadBpfRecorderArgsForCall []struct {
		arg1 *bpfrecorder.BpfRecorder
	}
	WriteFileStub        func(string, []byte, fs.FileMode) error
	writeFileMutex       sync.RWMutex
	writeFileArgsForCall []struct {
		arg1 string
		arg2 []byte
		arg3 fs.FileMode
	}
	writeFileReturns struct {
		result1 error
	}
	writeFileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImpl) CloseFile(arg1 *os.File) {
	fake.closeFileMutex.Lock()
	fake.closeFileArgsForCall = append(fake.closeFileArgsForCall, struct {
		arg1 *os.File
	}{arg1})
	stub := fake.CloseFileStub
	fake.recordInvocation("CloseFile", []interface{}{arg1})
	fake.closeFileMutex.Unlock()
	if stub != nil {
		fake.CloseFileStub(arg1)
	}
}

func (fake *FakeImpl) CloseFileCallCount() int {
	fake.closeFileMutex.RLock()
	defer fake.closeFileMutex.RUnlock()
	return len(fake.closeFileArgsForCall)
}

func (fake *FakeImpl) CloseFileCalls(stub func(*os.File)) {
	fake.closeFileMutex.Lock()
	defer fake.closeFileMutex.Unlock()
	fake.CloseFileStub = stub
}

func (fake *FakeImpl) CloseFileArgsForCall(i int) *os.File {
	fake.closeFileMutex.RLock()
	defer fake.closeFileMutex.RUnlock()
	argsForCall := fake.closeFileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CmdPid(arg1 *exec.Cmd) uint32 {
	fake.cmdPidMutex.Lock()
	ret, specificReturn := fake.cmdPidReturnsOnCall[len(fake.cmdPidArgsForCall)]
	fake.cmdPidArgsForCall = append(fake.cmdPidArgsForCall, struct {
		arg1 *exec.Cmd
	}{arg1})
	stub := fake.CmdPidStub
	fakeReturns := fake.cmdPidReturns
	fake.recordInvocation("CmdPid", []interface{}{arg1})
	fake.cmdPidMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) CmdPidCallCount() int {
	fake.cmdPidMutex.RLock()
	defer fake.cmdPidMutex.RUnlock()
	return len(fake.cmdPidArgsForCall)
}

func (fake *FakeImpl) CmdPidCalls(stub func(*exec.Cmd) uint32) {
	fake.cmdPidMutex.Lock()
	defer fake.cmdPidMutex.Unlock()
	fake.CmdPidStub = stub
}

func (fake *FakeImpl) CmdPidArgsForCall(i int) *exec.Cmd {
	fake.cmdPidMutex.RLock()
	defer fake.cmdPidMutex.RUnlock()
	argsForCall := fake.cmdPidArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CmdPidReturns(result1 uint32) {
	fake.cmdPidMutex.Lock()
	defer fake.cmdPidMutex.Unlock()
	fake.CmdPidStub = nil
	fake.cmdPidReturns = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeImpl) CmdPidReturnsOnCall(i int, result1 uint32) {
	fake.cmdPidMutex.Lock()
	defer fake.cmdPidMutex.Unlock()
	fake.CmdPidStub = nil
	if fake.cmdPidReturnsOnCall == nil {
		fake.cmdPidReturnsOnCall = make(map[int]struct {
			result1 uint32
		})
	}
	fake.cmdPidReturnsOnCall[i] = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeImpl) CmdStart(arg1 *exec.Cmd) error {
	fake.cmdStartMutex.Lock()
	ret, specificReturn := fake.cmdStartReturnsOnCall[len(fake.cmdStartArgsForCall)]
	fake.cmdStartArgsForCall = append(fake.cmdStartArgsForCall, struct {
		arg1 *exec.Cmd
	}{arg1})
	stub := fake.CmdStartStub
	fakeReturns := fake.cmdStartReturns
	fake.recordInvocation("CmdStart", []interface{}{arg1})
	fake.cmdStartMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) CmdStartCallCount() int {
	fake.cmdStartMutex.RLock()
	defer fake.cmdStartMutex.RUnlock()
	return len(fake.cmdStartArgsForCall)
}

func (fake *FakeImpl) CmdStartCalls(stub func(*exec.Cmd) error) {
	fake.cmdStartMutex.Lock()
	defer fake.cmdStartMutex.Unlock()
	fake.CmdStartStub = stub
}

func (fake *FakeImpl) CmdStartArgsForCall(i int) *exec.Cmd {
	fake.cmdStartMutex.RLock()
	defer fake.cmdStartMutex.RUnlock()
	argsForCall := fake.cmdStartArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CmdStartReturns(result1 error) {
	fake.cmdStartMutex.Lock()
	defer fake.cmdStartMutex.Unlock()
	fake.CmdStartStub = nil
	fake.cmdStartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) CmdStartReturnsOnCall(i int, result1 error) {
	fake.cmdStartMutex.Lock()
	defer fake.cmdStartMutex.Unlock()
	fake.CmdStartStub = nil
	if fake.cmdStartReturnsOnCall == nil {
		fake.cmdStartReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cmdStartReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) CmdWait(arg1 *exec.Cmd) error {
	fake.cmdWaitMutex.Lock()
	ret, specificReturn := fake.cmdWaitReturnsOnCall[len(fake.cmdWaitArgsForCall)]
	fake.cmdWaitArgsForCall = append(fake.cmdWaitArgsForCall, struct {
		arg1 *exec.Cmd
	}{arg1})
	stub := fake.CmdWaitStub
	fakeReturns := fake.cmdWaitReturns
	fake.recordInvocation("CmdWait", []interface{}{arg1})
	fake.cmdWaitMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) CmdWaitCallCount() int {
	fake.cmdWaitMutex.RLock()
	defer fake.cmdWaitMutex.RUnlock()
	return len(fake.cmdWaitArgsForCall)
}

func (fake *FakeImpl) CmdWaitCalls(stub func(*exec.Cmd) error) {
	fake.cmdWaitMutex.Lock()
	defer fake.cmdWaitMutex.Unlock()
	fake.CmdWaitStub = stub
}

func (fake *FakeImpl) CmdWaitArgsForCall(i int) *exec.Cmd {
	fake.cmdWaitMutex.RLock()
	defer fake.cmdWaitMutex.RUnlock()
	argsForCall := fake.cmdWaitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CmdWaitReturns(result1 error) {
	fake.cmdWaitMutex.Lock()
	defer fake.cmdWaitMutex.Unlock()
	fake.CmdWaitStub = nil
	fake.cmdWaitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) CmdWaitReturnsOnCall(i int, result1 error) {
	fake.cmdWaitMutex.Lock()
	defer fake.cmdWaitMutex.Unlock()
	fake.CmdWaitStub = nil
	if fake.cmdWaitReturnsOnCall == nil {
		fake.cmdWaitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cmdWaitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) Command(arg1 string, arg2 ...string) *exec.Cmd {
	fake.commandMutex.Lock()
	ret, specificReturn := fake.commandReturnsOnCall[len(fake.commandArgsForCall)]
	fake.commandArgsForCall = append(fake.commandArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2})
	stub := fake.CommandStub
	fakeReturns := fake.commandReturns
	fake.recordInvocation("Command", []interface{}{arg1, arg2})
	fake.commandMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) CommandCallCount() int {
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	return len(fake.commandArgsForCall)
}

func (fake *FakeImpl) CommandCalls(stub func(string, ...string) *exec.Cmd) {
	fake.commandMutex.Lock()
	defer fake.commandMutex.Unlock()
	fake.CommandStub = stub
}

func (fake *FakeImpl) CommandArgsForCall(i int) (string, []string) {
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	argsForCall := fake.commandArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) CommandReturns(result1 *exec.Cmd) {
	fake.commandMutex.Lock()
	defer fake.commandMutex.Unlock()
	fake.CommandStub = nil
	fake.commandReturns = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeImpl) CommandReturnsOnCall(i int, result1 *exec.Cmd) {
	fake.commandMutex.Lock()
	defer fake.commandMutex.Unlock()
	fake.CommandStub = nil
	if fake.commandReturnsOnCall == nil {
		fake.commandReturnsOnCall = make(map[int]struct {
			result1 *exec.Cmd
		})
	}
	fake.commandReturnsOnCall[i] = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeImpl) Create(arg1 string) (*os.File, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeImpl) CreateCalls(stub func(string) (*os.File, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeImpl) CreateArgsForCall(i int) string {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CreateReturns(result1 *os.File, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) CreateReturnsOnCall(i int, result1 *os.File, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *os.File
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GetName(arg1 seccomp.ScmpSyscall) (string, error) {
	fake.getNameMutex.Lock()
	ret, specificReturn := fake.getNameReturnsOnCall[len(fake.getNameArgsForCall)]
	fake.getNameArgsForCall = append(fake.getNameArgsForCall, struct {
		arg1 seccomp.ScmpSyscall
	}{arg1})
	stub := fake.GetNameStub
	fakeReturns := fake.getNameReturns
	fake.recordInvocation("GetName", []interface{}{arg1})
	fake.getNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) GetNameCallCount() int {
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	return len(fake.getNameArgsForCall)
}

func (fake *FakeImpl) GetNameCalls(stub func(seccomp.ScmpSyscall) (string, error)) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = stub
}

func (fake *FakeImpl) GetNameArgsForCall(i int) seccomp.ScmpSyscall {
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	argsForCall := fake.getNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) GetNameReturns(result1 string, result2 error) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = nil
	fake.getNameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GetNameReturnsOnCall(i int, result1 string, result2 error) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = nil
	if fake.getNameReturnsOnCall == nil {
		fake.getNameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getNameReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GoArchToSeccompArch(arg1 string) (seccompa.Arch, error) {
	fake.goArchToSeccompArchMutex.Lock()
	ret, specificReturn := fake.goArchToSeccompArchReturnsOnCall[len(fake.goArchToSeccompArchArgsForCall)]
	fake.goArchToSeccompArchArgsForCall = append(fake.goArchToSeccompArchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GoArchToSeccompArchStub
	fakeReturns := fake.goArchToSeccompArchReturns
	fake.recordInvocation("GoArchToSeccompArch", []interface{}{arg1})
	fake.goArchToSeccompArchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) GoArchToSeccompArchCallCount() int {
	fake.goArchToSeccompArchMutex.RLock()
	defer fake.goArchToSeccompArchMutex.RUnlock()
	return len(fake.goArchToSeccompArchArgsForCall)
}

func (fake *FakeImpl) GoArchToSeccompArchCalls(stub func(string) (seccompa.Arch, error)) {
	fake.goArchToSeccompArchMutex.Lock()
	defer fake.goArchToSeccompArchMutex.Unlock()
	fake.GoArchToSeccompArchStub = stub
}

func (fake *FakeImpl) GoArchToSeccompArchArgsForCall(i int) string {
	fake.goArchToSeccompArchMutex.RLock()
	defer fake.goArchToSeccompArchMutex.RUnlock()
	argsForCall := fake.goArchToSeccompArchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) GoArchToSeccompArchReturns(result1 seccompa.Arch, result2 error) {
	fake.goArchToSeccompArchMutex.Lock()
	defer fake.goArchToSeccompArchMutex.Unlock()
	fake.GoArchToSeccompArchStub = nil
	fake.goArchToSeccompArchReturns = struct {
		result1 seccompa.Arch
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GoArchToSeccompArchReturnsOnCall(i int, result1 seccompa.Arch, result2 error) {
	fake.goArchToSeccompArchMutex.Lock()
	defer fake.goArchToSeccompArchMutex.Unlock()
	fake.GoArchToSeccompArchStub = nil
	if fake.goArchToSeccompArchReturnsOnCall == nil {
		fake.goArchToSeccompArchReturnsOnCall = make(map[int]struct {
			result1 seccompa.Arch
			result2 error
		})
	}
	fake.goArchToSeccompArchReturnsOnCall[i] = struct {
		result1 seccompa.Arch
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) IteratorKey(arg1 *libbpfgo.BPFMapIterator) []byte {
	fake.iteratorKeyMutex.Lock()
	ret, specificReturn := fake.iteratorKeyReturnsOnCall[len(fake.iteratorKeyArgsForCall)]
	fake.iteratorKeyArgsForCall = append(fake.iteratorKeyArgsForCall, struct {
		arg1 *libbpfgo.BPFMapIterator
	}{arg1})
	stub := fake.IteratorKeyStub
	fakeReturns := fake.iteratorKeyReturns
	fake.recordInvocation("IteratorKey", []interface{}{arg1})
	fake.iteratorKeyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) IteratorKeyCallCount() int {
	fake.iteratorKeyMutex.RLock()
	defer fake.iteratorKeyMutex.RUnlock()
	return len(fake.iteratorKeyArgsForCall)
}

func (fake *FakeImpl) IteratorKeyCalls(stub func(*libbpfgo.BPFMapIterator) []byte) {
	fake.iteratorKeyMutex.Lock()
	defer fake.iteratorKeyMutex.Unlock()
	fake.IteratorKeyStub = stub
}

func (fake *FakeImpl) IteratorKeyArgsForCall(i int) *libbpfgo.BPFMapIterator {
	fake.iteratorKeyMutex.RLock()
	defer fake.iteratorKeyMutex.RUnlock()
	argsForCall := fake.iteratorKeyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) IteratorKeyReturns(result1 []byte) {
	fake.iteratorKeyMutex.Lock()
	defer fake.iteratorKeyMutex.Unlock()
	fake.IteratorKeyStub = nil
	fake.iteratorKeyReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeImpl) IteratorKeyReturnsOnCall(i int, result1 []byte) {
	fake.iteratorKeyMutex.Lock()
	defer fake.iteratorKeyMutex.Unlock()
	fake.IteratorKeyStub = nil
	if fake.iteratorKeyReturnsOnCall == nil {
		fake.iteratorKeyReturnsOnCall = make(map[int]struct {
			result1 []byte
		})
	}
	fake.iteratorKeyReturnsOnCall[i] = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeImpl) IteratorNext(arg1 *libbpfgo.BPFMapIterator) bool {
	fake.iteratorNextMutex.Lock()
	ret, specificReturn := fake.iteratorNextReturnsOnCall[len(fake.iteratorNextArgsForCall)]
	fake.iteratorNextArgsForCall = append(fake.iteratorNextArgsForCall, struct {
		arg1 *libbpfgo.BPFMapIterator
	}{arg1})
	stub := fake.IteratorNextStub
	fakeReturns := fake.iteratorNextReturns
	fake.recordInvocation("IteratorNext", []interface{}{arg1})
	fake.iteratorNextMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) IteratorNextCallCount() int {
	fake.iteratorNextMutex.RLock()
	defer fake.iteratorNextMutex.RUnlock()
	return len(fake.iteratorNextArgsForCall)
}

func (fake *FakeImpl) IteratorNextCalls(stub func(*libbpfgo.BPFMapIterator) bool) {
	fake.iteratorNextMutex.Lock()
	defer fake.iteratorNextMutex.Unlock()
	fake.IteratorNextStub = stub
}

func (fake *FakeImpl) IteratorNextArgsForCall(i int) *libbpfgo.BPFMapIterator {
	fake.iteratorNextMutex.RLock()
	defer fake.iteratorNextMutex.RUnlock()
	argsForCall := fake.iteratorNextArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) IteratorNextReturns(result1 bool) {
	fake.iteratorNextMutex.Lock()
	defer fake.iteratorNextMutex.Unlock()
	fake.IteratorNextStub = nil
	fake.iteratorNextReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeImpl) IteratorNextReturnsOnCall(i int, result1 bool) {
	fake.iteratorNextMutex.Lock()
	defer fake.iteratorNextMutex.Unlock()
	fake.IteratorNextStub = nil
	if fake.iteratorNextReturnsOnCall == nil {
		fake.iteratorNextReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.iteratorNextReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeImpl) LoadBpfRecorder(arg1 *bpfrecorder.BpfRecorder) error {
	fake.loadBpfRecorderMutex.Lock()
	ret, specificReturn := fake.loadBpfRecorderReturnsOnCall[len(fake.loadBpfRecorderArgsForCall)]
	fake.loadBpfRecorderArgsForCall = append(fake.loadBpfRecorderArgsForCall, struct {
		arg1 *bpfrecorder.BpfRecorder
	}{arg1})
	stub := fake.LoadBpfRecorderStub
	fakeReturns := fake.loadBpfRecorderReturns
	fake.recordInvocation("LoadBpfRecorder", []interface{}{arg1})
	fake.loadBpfRecorderMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) LoadBpfRecorderCallCount() int {
	fake.loadBpfRecorderMutex.RLock()
	defer fake.loadBpfRecorderMutex.RUnlock()
	return len(fake.loadBpfRecorderArgsForCall)
}

func (fake *FakeImpl) LoadBpfRecorderCalls(stub func(*bpfrecorder.BpfRecorder) error) {
	fake.loadBpfRecorderMutex.Lock()
	defer fake.loadBpfRecorderMutex.Unlock()
	fake.LoadBpfRecorderStub = stub
}

func (fake *FakeImpl) LoadBpfRecorderArgsForCall(i int) *bpfrecorder.BpfRecorder {
	fake.loadBpfRecorderMutex.RLock()
	defer fake.loadBpfRecorderMutex.RUnlock()
	argsForCall := fake.loadBpfRecorderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) LoadBpfRecorderReturns(result1 error) {
	fake.loadBpfRecorderMutex.Lock()
	defer fake.loadBpfRecorderMutex.Unlock()
	fake.LoadBpfRecorderStub = nil
	fake.loadBpfRecorderReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) LoadBpfRecorderReturnsOnCall(i int, result1 error) {
	fake.loadBpfRecorderMutex.Lock()
	defer fake.loadBpfRecorderMutex.Unlock()
	fake.LoadBpfRecorderStub = nil
	if fake.loadBpfRecorderReturnsOnCall == nil {
		fake.loadBpfRecorderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loadBpfRecorderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) MarshalIndent(arg1 any, arg2 string, arg3 string) ([]byte, error) {
	fake.marshalIndentMutex.Lock()
	ret, specificReturn := fake.marshalIndentReturnsOnCall[len(fake.marshalIndentArgsForCall)]
	fake.marshalIndentArgsForCall = append(fake.marshalIndentArgsForCall, struct {
		arg1 any
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.MarshalIndentStub
	fakeReturns := fake.marshalIndentReturns
	fake.recordInvocation("MarshalIndent", []interface{}{arg1, arg2, arg3})
	fake.marshalIndentMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) MarshalIndentCallCount() int {
	fake.marshalIndentMutex.RLock()
	defer fake.marshalIndentMutex.RUnlock()
	return len(fake.marshalIndentArgsForCall)
}

func (fake *FakeImpl) MarshalIndentCalls(stub func(any, string, string) ([]byte, error)) {
	fake.marshalIndentMutex.Lock()
	defer fake.marshalIndentMutex.Unlock()
	fake.MarshalIndentStub = stub
}

func (fake *FakeImpl) MarshalIndentArgsForCall(i int) (any, string, string) {
	fake.marshalIndentMutex.RLock()
	defer fake.marshalIndentMutex.RUnlock()
	argsForCall := fake.marshalIndentArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) MarshalIndentReturns(result1 []byte, result2 error) {
	fake.marshalIndentMutex.Lock()
	defer fake.marshalIndentMutex.Unlock()
	fake.MarshalIndentStub = nil
	fake.marshalIndentReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) MarshalIndentReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.marshalIndentMutex.Lock()
	defer fake.marshalIndentMutex.Unlock()
	fake.MarshalIndentStub = nil
	if fake.marshalIndentReturnsOnCall == nil {
		fake.marshalIndentReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.marshalIndentReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) PrintObj(arg1 printers.YAMLPrinter, arg2 runtime.Object, arg3 io.Writer) error {
	fake.printObjMutex.Lock()
	ret, specificReturn := fake.printObjReturnsOnCall[len(fake.printObjArgsForCall)]
	fake.printObjArgsForCall = append(fake.printObjArgsForCall, struct {
		arg1 printers.YAMLPrinter
		arg2 runtime.Object
		arg3 io.Writer
	}{arg1, arg2, arg3})
	stub := fake.PrintObjStub
	fakeReturns := fake.printObjReturns
	fake.recordInvocation("PrintObj", []interface{}{arg1, arg2, arg3})
	fake.printObjMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) PrintObjCallCount() int {
	fake.printObjMutex.RLock()
	defer fake.printObjMutex.RUnlock()
	return len(fake.printObjArgsForCall)
}

func (fake *FakeImpl) PrintObjCalls(stub func(printers.YAMLPrinter, runtime.Object, io.Writer) error) {
	fake.printObjMutex.Lock()
	defer fake.printObjMutex.Unlock()
	fake.PrintObjStub = stub
}

func (fake *FakeImpl) PrintObjArgsForCall(i int) (printers.YAMLPrinter, runtime.Object, io.Writer) {
	fake.printObjMutex.RLock()
	defer fake.printObjMutex.RUnlock()
	argsForCall := fake.printObjArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) PrintObjReturns(result1 error) {
	fake.printObjMutex.Lock()
	defer fake.printObjMutex.Unlock()
	fake.PrintObjStub = nil
	fake.printObjReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) PrintObjReturnsOnCall(i int, result1 error) {
	fake.printObjMutex.Lock()
	defer fake.printObjMutex.Unlock()
	fake.PrintObjStub = nil
	if fake.printObjReturnsOnCall == nil {
		fake.printObjReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.printObjReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) SyscallsGetValue(arg1 *bpfrecorder.BpfRecorder, arg2 uint32) ([]byte, error) {
	fake.syscallsGetValueMutex.Lock()
	ret, specificReturn := fake.syscallsGetValueReturnsOnCall[len(fake.syscallsGetValueArgsForCall)]
	fake.syscallsGetValueArgsForCall = append(fake.syscallsGetValueArgsForCall, struct {
		arg1 *bpfrecorder.BpfRecorder
		arg2 uint32
	}{arg1, arg2})
	stub := fake.SyscallsGetValueStub
	fakeReturns := fake.syscallsGetValueReturns
	fake.recordInvocation("SyscallsGetValue", []interface{}{arg1, arg2})
	fake.syscallsGetValueMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) SyscallsGetValueCallCount() int {
	fake.syscallsGetValueMutex.RLock()
	defer fake.syscallsGetValueMutex.RUnlock()
	return len(fake.syscallsGetValueArgsForCall)
}

func (fake *FakeImpl) SyscallsGetValueCalls(stub func(*bpfrecorder.BpfRecorder, uint32) ([]byte, error)) {
	fake.syscallsGetValueMutex.Lock()
	defer fake.syscallsGetValueMutex.Unlock()
	fake.SyscallsGetValueStub = stub
}

func (fake *FakeImpl) SyscallsGetValueArgsForCall(i int) (*bpfrecorder.BpfRecorder, uint32) {
	fake.syscallsGetValueMutex.RLock()
	defer fake.syscallsGetValueMutex.RUnlock()
	argsForCall := fake.syscallsGetValueArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) SyscallsGetValueReturns(result1 []byte, result2 error) {
	fake.syscallsGetValueMutex.Lock()
	defer fake.syscallsGetValueMutex.Unlock()
	fake.SyscallsGetValueStub = nil
	fake.syscallsGetValueReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) SyscallsGetValueReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.syscallsGetValueMutex.Lock()
	defer fake.syscallsGetValueMutex.Unlock()
	fake.SyscallsGetValueStub = nil
	if fake.syscallsGetValueReturnsOnCall == nil {
		fake.syscallsGetValueReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.syscallsGetValueReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) SyscallsIterator(arg1 *bpfrecorder.BpfRecorder) *libbpfgo.BPFMapIterator {
	fake.syscallsIteratorMutex.Lock()
	ret, specificReturn := fake.syscallsIteratorReturnsOnCall[len(fake.syscallsIteratorArgsForCall)]
	fake.syscallsIteratorArgsForCall = append(fake.syscallsIteratorArgsForCall, struct {
		arg1 *bpfrecorder.BpfRecorder
	}{arg1})
	stub := fake.SyscallsIteratorStub
	fakeReturns := fake.syscallsIteratorReturns
	fake.recordInvocation("SyscallsIterator", []interface{}{arg1})
	fake.syscallsIteratorMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) SyscallsIteratorCallCount() int {
	fake.syscallsIteratorMutex.RLock()
	defer fake.syscallsIteratorMutex.RUnlock()
	return len(fake.syscallsIteratorArgsForCall)
}

func (fake *FakeImpl) SyscallsIteratorCalls(stub func(*bpfrecorder.BpfRecorder) *libbpfgo.BPFMapIterator) {
	fake.syscallsIteratorMutex.Lock()
	defer fake.syscallsIteratorMutex.Unlock()
	fake.SyscallsIteratorStub = stub
}

func (fake *FakeImpl) SyscallsIteratorArgsForCall(i int) *bpfrecorder.BpfRecorder {
	fake.syscallsIteratorMutex.RLock()
	defer fake.syscallsIteratorMutex.RUnlock()
	argsForCall := fake.syscallsIteratorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) SyscallsIteratorReturns(result1 *libbpfgo.BPFMapIterator) {
	fake.syscallsIteratorMutex.Lock()
	defer fake.syscallsIteratorMutex.Unlock()
	fake.SyscallsIteratorStub = nil
	fake.syscallsIteratorReturns = struct {
		result1 *libbpfgo.BPFMapIterator
	}{result1}
}

func (fake *FakeImpl) SyscallsIteratorReturnsOnCall(i int, result1 *libbpfgo.BPFMapIterator) {
	fake.syscallsIteratorMutex.Lock()
	defer fake.syscallsIteratorMutex.Unlock()
	fake.SyscallsIteratorStub = nil
	if fake.syscallsIteratorReturnsOnCall == nil {
		fake.syscallsIteratorReturnsOnCall = make(map[int]struct {
			result1 *libbpfgo.BPFMapIterator
		})
	}
	fake.syscallsIteratorReturnsOnCall[i] = struct {
		result1 *libbpfgo.BPFMapIterator
	}{result1}
}

func (fake *FakeImpl) UnloadBpfRecorder(arg1 *bpfrecorder.BpfRecorder) {
	fake.unloadBpfRecorderMutex.Lock()
	fake.unloadBpfRecorderArgsForCall = append(fake.unloadBpfRecorderArgsForCall, struct {
		arg1 *bpfrecorder.BpfRecorder
	}{arg1})
	stub := fake.UnloadBpfRecorderStub
	fake.recordInvocation("UnloadBpfRecorder", []interface{}{arg1})
	fake.unloadBpfRecorderMutex.Unlock()
	if stub != nil {
		fake.UnloadBpfRecorderStub(arg1)
	}
}

func (fake *FakeImpl) UnloadBpfRecorderCallCount() int {
	fake.unloadBpfRecorderMutex.RLock()
	defer fake.unloadBpfRecorderMutex.RUnlock()
	return len(fake.unloadBpfRecorderArgsForCall)
}

func (fake *FakeImpl) UnloadBpfRecorderCalls(stub func(*bpfrecorder.BpfRecorder)) {
	fake.unloadBpfRecorderMutex.Lock()
	defer fake.unloadBpfRecorderMutex.Unlock()
	fake.UnloadBpfRecorderStub = stub
}

func (fake *FakeImpl) UnloadBpfRecorderArgsForCall(i int) *bpfrecorder.BpfRecorder {
	fake.unloadBpfRecorderMutex.RLock()
	defer fake.unloadBpfRecorderMutex.RUnlock()
	argsForCall := fake.unloadBpfRecorderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) WriteFile(arg1 string, arg2 []byte, arg3 fs.FileMode) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.writeFileMutex.Lock()
	ret, specificReturn := fake.writeFileReturnsOnCall[len(fake.writeFileArgsForCall)]
	fake.writeFileArgsForCall = append(fake.writeFileArgsForCall, struct {
		arg1 string
		arg2 []byte
		arg3 fs.FileMode
	}{arg1, arg2Copy, arg3})
	stub := fake.WriteFileStub
	fakeReturns := fake.writeFileReturns
	fake.recordInvocation("WriteFile", []interface{}{arg1, arg2Copy, arg3})
	fake.writeFileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) WriteFileCallCount() int {
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	return len(fake.writeFileArgsForCall)
}

func (fake *FakeImpl) WriteFileCalls(stub func(string, []byte, fs.FileMode) error) {
	fake.writeFileMutex.Lock()
	defer fake.writeFileMutex.Unlock()
	fake.WriteFileStub = stub
}

func (fake *FakeImpl) WriteFileArgsForCall(i int) (string, []byte, fs.FileMode) {
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	argsForCall := fake.writeFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) WriteFileReturns(result1 error) {
	fake.writeFileMutex.Lock()
	defer fake.writeFileMutex.Unlock()
	fake.WriteFileStub = nil
	fake.writeFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) WriteFileReturnsOnCall(i int, result1 error) {
	fake.writeFileMutex.Lock()
	defer fake.writeFileMutex.Unlock()
	fake.WriteFileStub = nil
	if fake.writeFileReturnsOnCall == nil {
		fake.writeFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeFileMutex.RLock()
	defer fake.closeFileMutex.RUnlock()
	fake.cmdPidMutex.RLock()
	defer fake.cmdPidMutex.RUnlock()
	fake.cmdStartMutex.RLock()
	defer fake.cmdStartMutex.RUnlock()
	fake.cmdWaitMutex.RLock()
	defer fake.cmdWaitMutex.RUnlock()
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	fake.goArchToSeccompArchMutex.RLock()
	defer fake.goArchToSeccompArchMutex.RUnlock()
	fake.iteratorKeyMutex.RLock()
	defer fake.iteratorKeyMutex.RUnlock()
	fake.iteratorNextMutex.RLock()
	defer fake.iteratorNextMutex.RUnlock()
	fake.loadBpfRecorderMutex.RLock()
	defer fake.loadBpfRecorderMutex.RUnlock()
	fake.marshalIndentMutex.RLock()
	defer fake.marshalIndentMutex.RUnlock()
	fake.printObjMutex.RLock()
	defer fake.printObjMutex.RUnlock()
	fake.syscallsGetValueMutex.RLock()
	defer fake.syscallsGetValueMutex.RUnlock()
	fake.syscallsIteratorMutex.RLock()
	defer fake.syscallsIteratorMutex.RUnlock()
	fake.unloadBpfRecorderMutex.RLock()
	defer fake.unloadBpfRecorderMutex.RUnlock()
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImpl) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}