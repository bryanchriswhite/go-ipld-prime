package gendemo

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _Map__String__Msg3 struct {
	m map[_String]*_Msg3
	t []_Map__String__Msg3__entry
}
type Map__String__Msg3 = *_Map__String__Msg3
type _Map__String__Msg3__entry struct {
	k _String
	v _Msg3
}

func (n *_Map__String__Msg3) LookupMaybe(k String) MaybeMsg3 {
	v, ok := n.m[*k]
	if !ok {
		return &_Map__String__Msg3__valueAbsent
	}
	return &_Msg3__Maybe{
		m: schema.Maybe_Value,
		v: v,
	}
}

var _Map__String__Msg3__valueAbsent = _Msg3__Maybe{m: schema.Maybe_Absent}

// TODO generate also a plain Lookup method that doesn't box and alloc if this type contains non-nullable values!
type _Map__String__Msg3__Maybe struct {
	m schema.Maybe
	v Map__String__Msg3
}
type MaybeMap__String__Msg3 = *_Map__String__Msg3__Maybe

func (m MaybeMap__String__Msg3) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMap__String__Msg3) IsUndefined() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMap__String__Msg3) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMap__String__Msg3) AsNode() ipld.Node {
	switch m.m {
	case schema.Maybe_Absent:
		return ipld.Undef
	case schema.Maybe_Null:
		return ipld.Null
	case schema.Maybe_Value:
		return m.v
	default:
		panic("unreachable")
	}
}
func (m MaybeMap__String__Msg3) Must() Map__String__Msg3 {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}

var _ ipld.Node = (Map__String__Msg3)(&_Map__String__Msg3{})
var _ schema.TypedNode = (Map__String__Msg3)(&_Map__String__Msg3{})

func (Map__String__Msg3) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n Map__String__Msg3) LookupString(k string) (ipld.Node, error) {
	var k2 _String
	if err := (_String__Style{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	v, exists := n.m[k2]
	if !exists {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(k)}
	}
	return v, nil
}
func (n Map__String__Msg3) Lookup(k ipld.Node) (ipld.Node, error) {
	k2, ok := k.(String)
	if !ok {
		panic("todo invalid key type error")
		// 'ipld.ErrInvalidKey{TypeName:"gendemo.Map__String__Msg3", Key:&_String{k}}' doesn't quite cut it: need room to explain the type, and it's not guaranteed k can be turned into a string at all
	}
	v, exists := n.m[*k2]
	if !exists {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(k2.String())}
	}
	return v, nil
}
func (Map__String__Msg3) LookupIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"gendemo.Map__String__Msg3"}.LookupIndex(0)
}
func (n Map__String__Msg3) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupString(seg.String())
}
func (n Map__String__Msg3) MapIterator() ipld.MapIterator {
	return &_Map__String__Msg3__MapItr{n, 0}
}

type _Map__String__Msg3__MapItr struct {
	n   Map__String__Msg3
	idx int
}

func (itr *_Map__String__Msg3__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= len(itr.n.t) {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	x := &itr.n.t[itr.idx]
	k = &x.k
	v = &x.v
	itr.idx++
	return
}
func (itr *_Map__String__Msg3__MapItr) Done() bool {
	return itr.idx >= len(itr.n.t)
}

func (Map__String__Msg3) ListIterator() ipld.ListIterator {
	return nil
}
func (n Map__String__Msg3) Length() int {
	return len(n.t)
}
func (Map__String__Msg3) IsUndefined() bool {
	return false
}
func (Map__String__Msg3) IsNull() bool {
	return false
}
func (Map__String__Msg3) AsBool() (bool, error) {
	return mixins.Map{"gendemo.Map__String__Msg3"}.AsBool()
}
func (Map__String__Msg3) AsInt() (int, error) {
	return mixins.Map{"gendemo.Map__String__Msg3"}.AsInt()
}
func (Map__String__Msg3) AsFloat() (float64, error) {
	return mixins.Map{"gendemo.Map__String__Msg3"}.AsFloat()
}
func (Map__String__Msg3) AsString() (string, error) {
	return mixins.Map{"gendemo.Map__String__Msg3"}.AsString()
}
func (Map__String__Msg3) AsBytes() ([]byte, error) {
	return mixins.Map{"gendemo.Map__String__Msg3"}.AsBytes()
}
func (Map__String__Msg3) AsLink() (ipld.Link, error) {
	return mixins.Map{"gendemo.Map__String__Msg3"}.AsLink()
}
func (Map__String__Msg3) Style() ipld.NodeStyle {
	return _Map__String__Msg3__Style{}
}

type _Map__String__Msg3__Style struct{}

func (_Map__String__Msg3__Style) NewBuilder() ipld.NodeBuilder {
	var nb _Map__String__Msg3__Builder
	nb.Reset()
	return &nb
}

type _Map__String__Msg3__Builder struct {
	_Map__String__Msg3__Assembler
}

func (nb *_Map__String__Msg3__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Map__String__Msg3__Builder) Reset() {
	var w _Map__String__Msg3
	var m schema.Maybe
	*nb = _Map__String__Msg3__Builder{_Map__String__Msg3__Assembler{w: &w, m: &m}}
}

type _Map__String__Msg3__Assembler struct {
	w     *_Map__String__Msg3
	m     *schema.Maybe
	state maState

	cm schema.Maybe
	ka _String__Assembler
	va _Msg3__Assembler
}

func (na *_Map__String__Msg3__Assembler) reset() {
	na.state = maState_initial
	na.ka.reset()
	na.va.reset()
}
func (na *_Map__String__Msg3__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if sizeHint < 0 {
		sizeHint = 0
	}
	if na.w == nil {
		na.w = &_Map__String__Msg3{}
	}
	na.w.m = make(map[_String]*_Msg3, sizeHint)
	na.w.t = make([]_Map__String__Msg3__entry, 0, sizeHint)
	return na, nil
}
func (_Map__String__Msg3__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3"}.BeginList(0)
}
func (na *_Map__String__Msg3__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"gendemo.Map__String__Msg3"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Map__String__Msg3__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3"}.AssignBool(false)
}
func (_Map__String__Msg3__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3"}.AssignInt(0)
}
func (_Map__String__Msg3__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3"}.AssignFloat(0)
}
func (_Map__String__Msg3__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3"}.AssignString("")
}
func (_Map__String__Msg3__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3"}.AssignBytes(nil)
}
func (_Map__String__Msg3__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3"}.AssignLink(nil)
}
func (na *_Map__String__Msg3__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Map__String__Msg3); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		case midvalue:
			panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v.ReprKind() != ipld.ReprKind_Map {
		return ipld.ErrWrongKind{TypeName: "gendemo.Map__String__Msg3", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
	}
	itr := v.MapIterator()
	for !itr.Done() {
		k, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleKey().AssignNode(k); err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_Map__String__Msg3__Assembler) Style() ipld.NodeStyle {
	return _Map__String__Msg3__Style{}
}
func (ma *_Map__String__Msg3__Assembler) keyFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.ka.w = nil
		tz := &ma.w.t[len(ma.w.t)-1]
		ma.cm = schema.Maybe_Absent
		ma.state = maState_expectValue
		ma.w.m[tz.k] = &tz.v
		ma.va.w = &tz.v
		ma.va.m = &ma.cm
		ma.ka.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__String__Msg3__Assembler) valueFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.va.w = nil
		ma.cm = schema.Maybe_Absent
		ma.state = maState_initial
		ma.va.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__String__Msg3__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleEntry cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleEntry cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleEntry cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleEntry cannot be called on an assembler that's already finished")
	}

	var k2 _String
	if err := (_String__Style{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	if _, exists := ma.w.m[k2]; exists {
		return nil, ipld.ErrRepeatedMapKey{&k2}
	}
	ma.w.t = append(ma.w.t, _Map__String__Msg3__entry{k: k2})
	tz := &ma.w.t[len(ma.w.t)-1]
	ma.state = maState_midValue

	ma.w.m[k2] = &tz.v
	ma.va.w = &tz.v
	ma.va.m = &ma.cm
	return &ma.va, nil
}
func (ma *_Map__String__Msg3__Assembler) AssembleKey() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleKey cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleKey cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleKey cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleKey cannot be called on an assembler that's already finished")
	}
	ma.w.t = append(ma.w.t, _Map__String__Msg3__entry{})
	ma.state = maState_midKey
	ma.ka.m = &ma.cm
	ma.ka.w = &ma.w.t[len(ma.w.t)-1].k
	return &ma.ka
}
func (ma *_Map__String__Msg3__Assembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		if !ma.keyFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
		} // if tidy success: carry on
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	return &ma.va
}
func (ma *_Map__String__Msg3__Assembler) Finish() error {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: Finish cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		panic("invalid state: Finish cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_Map__String__Msg3__Assembler) KeyStyle() ipld.NodeStyle {
	return _String__Style{}
}
func (ma *_Map__String__Msg3__Assembler) ValueStyle(_ string) ipld.NodeStyle {
	return _Msg3__Style{}
}
func (Map__String__Msg3) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n Map__String__Msg3) Representation() ipld.Node {
	return (*_Map__String__Msg3__Repr)(n)
}

type _Map__String__Msg3__Repr _Map__String__Msg3

var _ ipld.Node = &_Map__String__Msg3__Repr{}

func (_Map__String__Msg3__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (nr *_Map__String__Msg3__Repr) LookupString(k string) (ipld.Node, error) {
	v, err := (Map__String__Msg3)(nr).LookupString(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(Msg3).Representation(), nil
}
func (nr *_Map__String__Msg3__Repr) Lookup(k ipld.Node) (ipld.Node, error) {
	v, err := (Map__String__Msg3)(nr).Lookup(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(Msg3).Representation(), nil
}
func (_Map__String__Msg3__Repr) LookupIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"gendemo.Map__String__Msg3.Repr"}.LookupIndex(0)
}
func (n _Map__String__Msg3__Repr) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupString(seg.String())
}
func (nr *_Map__String__Msg3__Repr) MapIterator() ipld.MapIterator {
	return &_Map__String__Msg3__ReprMapItr{(Map__String__Msg3)(nr), 0}
}

type _Map__String__Msg3__ReprMapItr _Map__String__Msg3__MapItr

func (itr *_Map__String__Msg3__ReprMapItr) Next() (k ipld.Node, v ipld.Node, err error) {
	k, v, err = (*_Map__String__Msg3__MapItr)(itr).Next()
	if err != nil || v == ipld.Null {
		return
	}
	return k, v.(Msg3).Representation(), nil
}
func (itr *_Map__String__Msg3__ReprMapItr) Done() bool {
	return (*_Map__String__Msg3__MapItr)(itr).Done()
}

func (_Map__String__Msg3__Repr) ListIterator() ipld.ListIterator {
	return nil
}
func (rn *_Map__String__Msg3__Repr) Length() int {
	return len(rn.t)
}
func (_Map__String__Msg3__Repr) IsUndefined() bool {
	return false
}
func (_Map__String__Msg3__Repr) IsNull() bool {
	return false
}
func (_Map__String__Msg3__Repr) AsBool() (bool, error) {
	return mixins.Map{"gendemo.Map__String__Msg3.Repr"}.AsBool()
}
func (_Map__String__Msg3__Repr) AsInt() (int, error) {
	return mixins.Map{"gendemo.Map__String__Msg3.Repr"}.AsInt()
}
func (_Map__String__Msg3__Repr) AsFloat() (float64, error) {
	return mixins.Map{"gendemo.Map__String__Msg3.Repr"}.AsFloat()
}
func (_Map__String__Msg3__Repr) AsString() (string, error) {
	return mixins.Map{"gendemo.Map__String__Msg3.Repr"}.AsString()
}
func (_Map__String__Msg3__Repr) AsBytes() ([]byte, error) {
	return mixins.Map{"gendemo.Map__String__Msg3.Repr"}.AsBytes()
}
func (_Map__String__Msg3__Repr) AsLink() (ipld.Link, error) {
	return mixins.Map{"gendemo.Map__String__Msg3.Repr"}.AsLink()
}
func (_Map__String__Msg3__Repr) Style() ipld.NodeStyle {
	return _Map__String__Msg3__ReprStyle{}
}

type _Map__String__Msg3__ReprStyle struct{}

func (_Map__String__Msg3__ReprStyle) NewBuilder() ipld.NodeBuilder {
	var nb _Map__String__Msg3__ReprBuilder
	nb.Reset()
	return &nb
}

type _Map__String__Msg3__ReprBuilder struct {
	_Map__String__Msg3__ReprAssembler
}

func (nb *_Map__String__Msg3__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Map__String__Msg3__ReprBuilder) Reset() {
	var w _Map__String__Msg3
	var m schema.Maybe
	*nb = _Map__String__Msg3__ReprBuilder{_Map__String__Msg3__ReprAssembler{w: &w, m: &m}}
}

type _Map__String__Msg3__ReprAssembler struct {
	w     *_Map__String__Msg3
	m     *schema.Maybe
	state maState

	cm schema.Maybe
	ka _String__ReprAssembler
	va _Msg3__ReprAssembler
}

func (na *_Map__String__Msg3__ReprAssembler) reset() {
	na.state = maState_initial
	na.ka.reset()
	na.va.reset()
}
func (na *_Map__String__Msg3__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if sizeHint < 0 {
		sizeHint = 0
	}
	if na.w == nil {
		na.w = &_Map__String__Msg3{}
	}
	na.w.m = make(map[_String]*_Msg3, sizeHint)
	na.w.t = make([]_Map__String__Msg3__entry, 0, sizeHint)
	return na, nil
}
func (_Map__String__Msg3__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3.Repr"}.BeginList(0)
}
func (na *_Map__String__Msg3__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"gendemo.Map__String__Msg3.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Map__String__Msg3__ReprAssembler) AssignBool(bool) error {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3.Repr"}.AssignBool(false)
}
func (_Map__String__Msg3__ReprAssembler) AssignInt(int) error {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3.Repr"}.AssignInt(0)
}
func (_Map__String__Msg3__ReprAssembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3.Repr"}.AssignFloat(0)
}
func (_Map__String__Msg3__ReprAssembler) AssignString(string) error {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3.Repr"}.AssignString("")
}
func (_Map__String__Msg3__ReprAssembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3.Repr"}.AssignBytes(nil)
}
func (_Map__String__Msg3__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"gendemo.Map__String__Msg3.Repr"}.AssignLink(nil)
}
func (na *_Map__String__Msg3__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Map__String__Msg3); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		case midvalue:
			panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v.ReprKind() != ipld.ReprKind_Map {
		return ipld.ErrWrongKind{TypeName: "gendemo.Map__String__Msg3.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
	}
	itr := v.MapIterator()
	for !itr.Done() {
		k, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleKey().AssignNode(k); err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_Map__String__Msg3__ReprAssembler) Style() ipld.NodeStyle {
	return _Map__String__Msg3__ReprStyle{}
}
func (ma *_Map__String__Msg3__ReprAssembler) keyFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.ka.w = nil
		tz := &ma.w.t[len(ma.w.t)-1]
		ma.cm = schema.Maybe_Absent
		ma.state = maState_expectValue
		ma.w.m[tz.k] = &tz.v
		ma.va.w = &tz.v
		ma.va.m = &ma.cm
		ma.ka.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__String__Msg3__ReprAssembler) valueFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.va.w = nil
		ma.cm = schema.Maybe_Absent
		ma.state = maState_initial
		ma.va.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__String__Msg3__ReprAssembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleEntry cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleEntry cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleEntry cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleEntry cannot be called on an assembler that's already finished")
	}

	var k2 _String
	if err := (_String__ReprStyle{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	if _, exists := ma.w.m[k2]; exists {
		return nil, ipld.ErrRepeatedMapKey{&k2}
	}
	ma.w.t = append(ma.w.t, _Map__String__Msg3__entry{k: k2})
	tz := &ma.w.t[len(ma.w.t)-1]
	ma.state = maState_midValue

	ma.w.m[k2] = &tz.v
	ma.va.w = &tz.v
	ma.va.m = &ma.cm
	return &ma.va, nil
}
func (ma *_Map__String__Msg3__ReprAssembler) AssembleKey() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleKey cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleKey cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleKey cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleKey cannot be called on an assembler that's already finished")
	}
	ma.w.t = append(ma.w.t, _Map__String__Msg3__entry{})
	ma.state = maState_midKey
	ma.ka.m = &ma.cm
	ma.ka.w = &ma.w.t[len(ma.w.t)-1].k
	return &ma.ka
}
func (ma *_Map__String__Msg3__ReprAssembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		if !ma.keyFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
		} // if tidy success: carry on
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	return &ma.va
}
func (ma *_Map__String__Msg3__ReprAssembler) Finish() error {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: Finish cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		panic("invalid state: Finish cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_Map__String__Msg3__ReprAssembler) KeyStyle() ipld.NodeStyle {
	return _String__ReprStyle{}
}
func (ma *_Map__String__Msg3__ReprAssembler) ValueStyle(_ string) ipld.NodeStyle {
	return _Msg3__ReprStyle{}
}